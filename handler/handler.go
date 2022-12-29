package handler

import (
	"html/template"
	"net/http"

	"github.com/1ort/goimbo/model"
	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

type WebConfig struct {
	R             *gin.Engine //router
	BaseURL       string
	Userspace     model.Userspace
	CookieSecret  string
	XCSRFSecret   string
	EnableCaptcha bool
}

type APIConfig struct {
	R         *gin.Engine //router
	BaseURL   string
	Userspace model.Userspace
}

type WebHandler struct {
	userspace     model.Userspace
	r             *gin.Engine
	enableCaptcha bool
}

type APIHandler struct {
	userspace model.Userspace
	r         *gin.Engine
}

func SetWebHandler(cfg *WebConfig) {
	h := &WebHandler{
		userspace:     cfg.Userspace,
		r:             cfg.R,
		enableCaptcha: cfg.EnableCaptcha,
	}

	funcmap := template.FuncMap{
		"intRange":   IntRange,
		"formatBody": FormatBody,
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
	cfg.R.SetHTMLTemplate(tmpl)
	cfg.R.StaticFile("/styles.css", "./res/static/styles.css")
	cfg.R.StaticFile("/favicon.ico", "./res/static/favicon.ico")
	cfg.R.StaticFile("error_img.png", "./res/static/gopher_vojak_transparent.png")

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
	if cfg.EnableCaptcha {
		web.GET("/captcha/*id", gin.WrapH(captcha.Server(240, 80)))
	}

}

func SetAPIHandler(cfg *APIConfig) {
	h := &APIHandler{
		userspace: cfg.Userspace,
		r:         cfg.R,
	}
	api := cfg.R.Group(cfg.BaseURL)
	api.GET("/boards", h.getBoards)

	apiBoard := api.Group("/:board")
	apiBoard.GET("/", h.getBoard)
	apiBoard.GET("/:page", h.getPage)
	apiBoard.GET("/thread/:op", h.getThread)
}
