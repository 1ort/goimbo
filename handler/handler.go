package handler

import (
	"html/template"

	"github.com/1ort/goimbo/model"
	"github.com/gin-gonic/gin"
)

// Сюда будем передавать всё что нужно для инициализации хандлера
type Config struct {
	R         *gin.Engine //router
	BaseURL   string
	Userspace model.Userspace
}

type WebHandler struct {
	userspace model.Userspace
	r         *gin.Engine
}

type APIHandler struct {
	userspace model.Userspace
	r         *gin.Engine
}

func SetWebHandler(cfg *Config) {
	h := &WebHandler{
		userspace: cfg.Userspace,
		r:         cfg.R,
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
		))
	cfg.R.SetHTMLTemplate(tmpl)
	cfg.R.StaticFile("/styles.css", "./res/static/styles.css")
	cfg.R.StaticFile("/favicon.ico", "./res/static/favicon.ico")

	web := cfg.R.Group(cfg.BaseURL)
	web.GET("/", h.mainPage)

	webBoard := web.Group("/:board")
	webBoard.POST("/newthread", h.newthread)
	webBoard.GET("/", h.redirectToZeroPage) //redirects to /page/0/
	webBoard.GET("/page/:page", h.boardPage)

	webThread := webBoard.Group("/thread/:thread")
	webThread.GET("/", h.threadPage)
	webThread.POST("/reply", h.reply)
}

func SetAPIHandler(cfg *Config) {
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
