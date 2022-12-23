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
}

type ApiHandler struct {
	userspace model.Userspace
}

func SetWebHandler(cfg *HandlerConfig) {
	h := &WebHandler{
		userspace: cfg.Userspace,
	}

	tmpl := template.Must(template.ParseFiles("./res/templates/dir.html"))
	cfg.R.SetHTMLTemplate(tmpl)

	web := cfg.R.Group(cfg.BaseUrl)
	web.GET("/", h.main_page)

	web_board := web.Group("/:board")
	web_board.GET("/", h.board_page) //TODO: redirect to board index page. Separate thread pages
	web_board.POST("/newthread", h.newthread)

	web_thread := web_board.Group("/thread/:thread")
	web_thread.GET("/", h.thread_page)
	web_thread.POST("/reply", h.reply)
}

func SetApiHandler(cfg *HandlerConfig) {
	h := &ApiHandler{
		userspace: cfg.Userspace,
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
