package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/1ort/goimbo/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

type WebConfig struct {
	R            *gin.Engine //router
	BaseURL      string
	Userspace    model.Userspace
	CookieSecret string
	XCSRFSecret  string
	Captcha      WebCaptchaWrapper
}

type WebHandler struct {
	userspace model.Userspace
	r         *gin.Engine
	captcha   WebCaptchaWrapper
}

func SetWebHandler(cfg *WebConfig) {
	h := &WebHandler{
		userspace: cfg.Userspace,
		r:         cfg.R,
		captcha:   cfg.Captcha,
	}

	funcmap := template.FuncMap{
		"intRange":   IntRange,
		"formatBody": FormatBody,
		"captcha":    cfg.Captcha.presentCaptcha,
	}

	tmpl := template.Must(
		template.New("").Funcs(funcmap).ParseFiles(
			"res/templates/post.partial.tmpl",
			"res/templates/postform.partial.tmpl",
			"res/templates/thread_preview.partial.tmpl",
			"res/templates/footer.partial.tmpl",
			"res/templates/navbar.partial.tmpl",

			"res/templates/board.page.tmpl",
			"res/templates/thread.page.tmpl",
			"res/templates/dir.page.tmpl",
			"res/templates/error.page.tmpl",
		))
	h.r.SetHTMLTemplate(tmpl)
	h.r.StaticFile("/favicon.ico", "./res/static/favicon.ico")
	h.r.Static("/static", "./res/static/")

	web := cfg.R.Group(cfg.BaseURL)
	store := cookie.NewStore([]byte(cfg.CookieSecret))
	web.Use(sessions.Sessions("main", store))
	web.Use(csrf.Middleware(csrf.Options{
		Secret: cfg.XCSRFSecret,
		ErrorFunc: func(c *gin.Context) {
			c.Redirect(http.StatusSeeOther, "/") //TODO: error page template
			c.Abort()
		},
	}))
	web.GET("/", h.mainPage)

	webBoard := web.Group("/:board")
	webBoard.POST("/newthread", h.newthread)
	webBoard.GET("/", h.redirectToZeroPage) //redirects to /page/0/
	webBoard.GET("/page/:page", h.boardPage)

	webThread := webBoard.Group("/thread/:thread")
	webThread.GET("/", h.threadPage)
	webThread.POST("/reply", h.reply)
	web.GET("/captcha/:id", gin.WrapF(h.captcha.ServeHTTP))

}

func (h *WebHandler) handleError(c *gin.Context, err error) bool {
	if err == nil {
		return false
	}
	m := gin.H{
		"status":  model.Status(err),
		"message": err,
	}
	if model.Status(err) == 500 {
		m["message"] = "Something went wrong"
	}
	c.HTML(model.Status(err), "error.page.tmpl", m)
	return true
}

func (h *WebHandler) mainPage(c *gin.Context) {
	boards, err := h.userspace.GetBoards(c.Request.Context())
	if handled := h.handleError(c, err); handled {
		return
	}
	c.HTML(http.StatusOK, "dir.page.tmpl", gin.H{
		"Boards":      boards,
		"XCSRF_TOKEN": csrf.GetToken(c),
	})
}

func (h *WebHandler) redirectToZeroPage(c *gin.Context) {
	c.Redirect(http.StatusFound, fmt.Sprintf("/%s/page/0", c.Param("board")))
}

func (h *WebHandler) boardPage(c *gin.Context) {
	var p BoardPageRequest
	if err := c.ShouldBindUri(&p); err != nil {
		h.handleError(c, model.NewBadRequest("Invalid board or page"))
		return
	}
	boardData, err := h.userspace.GetBoard(c.Request.Context(), p.Board)
	if handled := h.handleError(c, err); handled {
		return
	}
	pageData, err := h.userspace.GetBoardPage(c.Request.Context(), p.Board, p.Page)
	if handled := h.handleError(c, err); handled {
		return
	}
	c.HTML(http.StatusOK, "board.page.tmpl", gin.H{
		"Page":        pageData.Page,
		"Threads":     pageData.Threads,
		"Board":       boardData,
		"XCSRF_TOKEN": csrf.GetToken(c),
	})
}

func (h *WebHandler) threadPage(c *gin.Context) {
	var p ThreadPageRequest
	if err := c.ShouldBindUri(&p); err != nil {
		h.handleError(c, model.NewBadRequest("Invalid thread"))
		return
	}
	boardData, err := h.userspace.GetBoard(c.Request.Context(), p.Board)
	if handled := h.handleError(c, err); handled {
		return
	}
	threadData, err := h.userspace.GetThread(c.Request.Context(), p.Board, p.Thread)
	if handled := h.handleError(c, err); handled {
		return
	}
	xCSRFToken := csrf.GetToken(c)
	c.HTML(http.StatusOK, "thread.page.tmpl", gin.H{
		"board_data":  boardData,
		"OP":          threadData.OP,
		"Replies":     threadData.Replies,
		"XCSRF_TOKEN": xCSRFToken,
	})
}

func (h *WebHandler) reply(c *gin.Context) {
	var t ThreadPageRequest
	var p PostRequest

	if err := c.ShouldBindUri(&t); err != nil {
		h.handleError(c, model.NewBadRequest("Invalid thread"))
		return
	}
	if err := c.ShouldBind(&p); err != nil {
		h.handleError(c, model.NewBadRequest("Invalid post content"))
		return
	}

	err := h.captcha.verify(c)
	if handled := h.handleError(c, err); handled {
		return
	}

	//files uploading
	form, err := c.MultipartForm()
	if err != nil {
		h.handleError(c, model.NewBadRequest("Form error"))
		return
	}
	files := form.File["files"]
	for _, file := range files {
		filename := filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file, filename); err != nil {
			h.handleError(c, model.NewBadRequest("File upload error"))
			return
		}
	}
	//end file uploading

	newPost, err := h.userspace.Reply(c.Request.Context(), t.Board, p.Com, t.Thread)
	if handled := h.handleError(c, err); handled {
		return
	}
	c.Redirect(http.StatusFound, fmt.Sprintf("/%s/thread/%v#%v", t.Board, t.Thread, newPost.No))
}

func (h *WebHandler) newthread(c *gin.Context) {
	var b BoardRequest
	var p PostRequest

	if err := c.ShouldBindUri(&b); err != nil {
		h.handleError(c, model.NewBadRequest("Invalid board"))
		return
	}
	if err := c.ShouldBind(&p); err != nil {
		h.handleError(c, model.NewBadRequest("Invalid post content"))
		return
	}

	err := h.captcha.verify(c)
	if handled := h.handleError(c, err); handled {
		return
	}

	newPost, err := h.userspace.NewThread(c.Request.Context(), b.Board, p.Com)
	if handled := h.handleError(c, err); handled {
		return
	}
	c.Redirect(http.StatusFound, fmt.Sprintf("/%s/thread/%v", b.Board, newPost.No))
}
