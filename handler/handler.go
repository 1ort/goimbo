package handler

import (
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
	R         *gin.Engine //router
	BaseUrl   string
	Userspace model.Userspace
}

func NewHandler(cfg *HandlerConfig) {
	h := &Handler{
		userspace: cfg.Userspace,
	}
	api := cfg.R.Group(cfg.BaseUrl)

	api.GET("/boards", h.get_boards)
	board := api.Group("/:board")
	board.GET("/threads", h.get_threads)
	board.GET("/catalog", h.get_catalog)
	board.GET("/archive", h.get_archive)
	board.GET("/:page", h.get_page)
	board.GET("/thread/:op", h.get_thread)

}
