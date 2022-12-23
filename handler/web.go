package handler

/*
	web := cfg.R.Group(cfg.WebBaseUrl)
	web.GET("/", h.main_page)
	web_board := web.Group("/:board")
	web_board.GET("/", h.board_page)
	web_thread := web_board.Group("/:thread")
	web_thread.GET("/", h.thread_page)
	web_board.POST("/reply", h.reply)
	web_board.POST("/newthread", h.newthread)
*/

import (
	"net/http"

	"github.com/1ort/goimbo/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) main_page(c *gin.Context) {
	boards, err := h.userspace.Boards(c.Request.Context())
	if err != nil {
		c.JSON(model.Status(err), gin.H{
			"status": model.Status(err),
			"result": err,
		})
		return
	}
	c.HTML(http.StatusOK, "dir.html", gin.H{
		"Boards": boards,
	})
}

func (h *Handler) board_page(c *gin.Context) {

}

func (h *Handler) thread_page(c *gin.Context) {

}

func (h *Handler) reply(c *gin.Context) {

}

func (h *Handler) newthread(c *gin.Context) {

}
