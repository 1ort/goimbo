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
	"fmt"
	"net/http"
	"strconv"

	"github.com/1ort/goimbo/model"
	"github.com/gin-gonic/gin"
)

func (h *WebHandler) main_page(c *gin.Context) {
	boards, err := h.userspace.GetBoards(c.Request.Context())
	if err != nil {
		c.JSON(model.Status(err), gin.H{
			"status": model.Status(err),
			"result": err,
		})
		return
	}
	c.HTML(http.StatusOK, "dir.page.tmpl", gin.H{
		"Boards": boards,
	})
}

func (h *WebHandler) redirect_to_zero_page(c *gin.Context) {
	board := c.Param("board")
	//c.Params = append(c.Params, gin.Param{Key: "page", Value: "0"})
	//h.board_page(c)
	c.Redirect(http.StatusFound, fmt.Sprintf("/%s/page/0", board))
}

func (h *WebHandler) board_page(c *gin.Context) {
	board := c.Param("board")
	page := c.Param("page")
	page_n, err := strconv.Atoi(page)
	if err != nil {
		err := model.NewNotFound("page", page)
		c.JSON(model.Status(err), gin.H{
			"status": model.Status(err),
			"result": err,
		})
		return
	}
	board_data, err := h.userspace.GetBoard(c.Request.Context(), board)
	if err != nil {
		c.JSON(model.Status(err), gin.H{
			"status": model.Status(err),
			"result": err,
		})
		return
	}
	page_data, err := h.userspace.GetBoardPage(c.Request.Context(), board, page_n)
	if err != nil {
		c.JSON(model.Status(err), gin.H{
			"status": model.Status(err),
			"result": err,
		})
		return
	}
	c.HTML(http.StatusOK, "board.page.tmpl", gin.H{
		"Page":    page_data.Page,
		"Threads": page_data.Threads,
		"Board":   board_data,
	})
}

func (h *WebHandler) thread_page(c *gin.Context) {
	board := c.Param("board")
	thread := c.Param("thread")
	board_data, err := h.userspace.GetBoard(c.Request.Context(), board)
	if err != nil {
		c.JSON(model.Status(err), gin.H{
			"status": model.Status(err),
			"result": err,
		})
		return
	}
	thread_n, err := strconv.Atoi(thread)
	if err != nil {
		err := model.NewNotFound("page", thread)
		c.JSON(model.Status(err), gin.H{
			"status": model.Status(err),
			"result": err,
		})
		return
	}
	thread_data, err := h.userspace.GetThread(c.Request.Context(), board, thread_n)
	if err != nil {
		err := model.NewNotFound("page", thread)
		c.JSON(model.Status(err), gin.H{
			"status": model.Status(err),
			"result": err,
		})
		return
	}
	c.HTML(http.StatusOK, "thread.page.tmpl", gin.H{
		"board_data": board_data,
		"OP":         thread_data.OP,
		"Replies":    thread_data.Replies,
	})
}

func (h *WebHandler) reply(c *gin.Context) {

}

func (h *WebHandler) newthread(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"res": "newthread",
	})
}
