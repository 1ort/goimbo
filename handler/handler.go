package handler

import (
	"html/template"

	"github.com/1ort/goimbo/model"
	"github.com/gin-gonic/gin"
)

// Сюда будем передавать всё что нужно для инициализации хандлера
type HandlerConfig struct {
	R         *gin.Engine //router
	BaseUrl   string
	Userspace model.Userspace
}

type WebHandler struct {
	userspace model.Userspace
	r         *gin.Engine
}

type ApiHandler struct {
	userspace model.Userspace
	r         *gin.Engine
}

func SetWebHandler(cfg *HandlerConfig) {
	h := &WebHandler{
		userspace: cfg.Userspace,
		r:         cfg.R,
	}

	funcmap := template.FuncMap{
		"intRange": IntRange,
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
		))
	cfg.R.SetHTMLTemplate(tmpl)
	cfg.R.StaticFile("/styles.css", "./res/static/styles.css")
	cfg.R.StaticFile("/favicon.ico", "./res/static/favicon.ico")

	web := cfg.R.Group(cfg.BaseUrl)
	web.GET("/", h.main_page)

	web_board := web.Group("/:board")
	web_board.POST("/newthread", h.newthread)
	web_board.GET("/", h.redirect_to_zero_page) //TODO: redirect to /page/0/
	web_board.GET("/page/:page", h.board_page)

	web_thread := web_board.Group("/thread/:thread")
	web_thread.GET("/", h.thread_page)
	web_thread.POST("/reply", h.reply)
}

func SetApiHandler(cfg *HandlerConfig) {
	h := &ApiHandler{
		userspace: cfg.Userspace,
		r:         cfg.R,
	}
	api := cfg.R.Group(cfg.BaseUrl)
	api.GET("/boards", h.get_boards)
	api_board := api.Group("/:board")

	api_board.GET("/threads", h.get_threads)
	api_board.GET("/catalog", h.get_catalog)
	api_board.GET("/archive", h.get_archive)
	api_board.GET("/:page", h.get_page)
	api_board.GET("/thread/:op", h.get_thread)
}
