package handler

import (
	"html/template"

	"github.com/1ort/goimbo/model"
	"github.com/gin-gonic/gin"
)

// Будет содержать всякие штуки типа коннектов к ДБ и хранилищ куки
type Handler struct {
	userspace model.Userspace
	//postRepo  model.PostRepository
}

// Сюда будем передавать всё что нужно для инициализации хандлера
type HandlerConfig struct {
	R          *gin.Engine //router
	ApiBaseUrl string
	WebBaseUrl string
	Userspace  model.Userspace
}

func NewHandler(cfg *HandlerConfig) {
	h := &Handler{
		userspace: cfg.Userspace,
	}

	tmpl := template.Must(template.ParseFiles("./res/templates/dir.html"))
	cfg.R.SetHTMLTemplate(tmpl)

	api := cfg.R.Group(cfg.ApiBaseUrl)
	api.GET("/boards", h.get_boards)
	api_board := api.Group("/:board")
	api_board.GET("/threads", h.get_threads)
	api_board.GET("/catalog", h.get_catalog)
	api_board.GET("/archive", h.get_archive)
	api_board.GET("/:page", h.get_page)
	api_board.GET("/thread/:op", h.get_thread)

	web := cfg.R.Group(cfg.WebBaseUrl)
	web.GET("/", h.main_page)
	web_board := web.Group("/:board")
	web_board.GET("/", h.board_page)
	web_thread := web_board.Group("/:thread")
	web_thread.GET("/", h.thread_page)
	web_board.POST("/reply", h.reply)
	web_board.POST("/newthread", h.newthread)
}
