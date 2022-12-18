package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Будет содержать всякие штуки типа коннектов к ДБ и хранилищ куки
type Handler struct{}

// Сюда будем передавать всё что нужно для инициализации хандлера
type HandlerConfig struct {
	R       *gin.Engine //router
	BaseUrl string
}

func NewHandler(cfg *HandlerConfig) {
	h := &Handler{}
	api := cfg.R.Group(cfg.BaseUrl)

	board := api.Group("/:board")

	board.GET("/threads", h.get_threads)
	board.GET("/catalog", h.get_catalog)
	board.GET("/archive", h.get_archive)
	board.GET("/:page", h.get_page)
	board.GET("/thread/:op", h.get_thread)

	//board.POST("/newpost", h.NewPost)
	//api.POST("/newboard", h.NewBoard)

}

func (h *Handler) get_boards(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"result": "board list here",
	})
}

func (h *Handler) get_threads(c *gin.Context) {
	board := c.Param("board")
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"result": fmt.Sprintf("Thread list on %s board here", board),
	})

}

func (h *Handler) get_catalog(c *gin.Context) {
	board := c.Param("board")
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"result": fmt.Sprintf("Catalog of %s board here", board),
	})

}

func (h *Handler) get_archive(c *gin.Context) {
	board := c.Param("board")
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"result": fmt.Sprintf("Archive of %s board here", board),
	})
}

func (h *Handler) get_page(c *gin.Context) {
	board := c.Param("board")
	page := c.Param("page")
	if _, err := strconv.Atoi(page); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"result": fmt.Sprintf("Page number must be an integer, got %s instead", page),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"result": fmt.Sprintf("Page number %s on %s board here", page, board),
		})
	}
}

func (h *Handler) get_thread(c *gin.Context) {
	board := c.Param("board")
	op := c.Param("op")
	if _, err := strconv.Atoi(op); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"result": fmt.Sprintf("Thread number must be an integer, got %s instead", op),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"result": fmt.Sprintf("Thread number %s content on %s board here", op, board),
		})
	}
}
