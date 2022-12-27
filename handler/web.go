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

func (h *WebHandler) mainPage(c *gin.Context) {
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

func (h *WebHandler) redirectToZeroPage(c *gin.Context) {
	board := c.Param("board")
	//c.Params = append(c.Params, gin.Param{Key: "page", Value: "0"})
	//h.board_page(c)
	c.Redirect(http.StatusFound, fmt.Sprintf("/%s/page/0", board))
}

func (h *WebHandler) boardPage(c *gin.Context) {
	board := c.Param("board")
	rawPage := c.Param("page")
	page, err := strconv.Atoi(rawPage)
	if err != nil {
		err := model.NewNotFound("page", rawPage)
		c.JSON(model.Status(err), gin.H{
			"status": model.Status(err),
			"result": err,
		})
		return
	}
	boardData, err := h.userspace.GetBoard(c.Request.Context(), board)
	if err != nil {
		c.JSON(model.Status(err), gin.H{
			"status": model.Status(err),
			"result": err,
		})
		return
	}
	pageData, err := h.userspace.GetBoardPage(c.Request.Context(), board, page)
	if err != nil {
		c.JSON(model.Status(err), gin.H{
			"status": model.Status(err),
			"result": err,
		})
		return
	}
	c.HTML(http.StatusOK, "board.page.tmpl", gin.H{
		"Page":    pageData.Page,
		"Threads": pageData.Threads,
		"Board":   boardData,
	})
}

func (h *WebHandler) threadPage(c *gin.Context) {
	board := c.Param("board")
	rawThread := c.Param("thread")
	boardData, err := h.userspace.GetBoard(c.Request.Context(), board)
	if err != nil {
		c.JSON(model.Status(err), gin.H{
			"status": model.Status(err),
			"result": err,
		})
		return
	}
	thread, err := strconv.Atoi(rawThread)
	if err != nil {
		err := model.NewNotFound("thread", rawThread)
		c.JSON(model.Status(err), gin.H{
			"status": model.Status(err),
			"result": err,
		})
		return
	}
	threadData, err := h.userspace.GetThread(c.Request.Context(), board, thread)
	if err != nil {
		fmt.Printf("%v \n", err)
		err := model.NewNotFound("page", rawThread)
		c.JSON(model.Status(err), gin.H{
			"status": model.Status(err),
			"result": err,
		})
		return
	}
	c.HTML(http.StatusOK, "thread.page.tmpl", gin.H{
		"board_data": boardData,
		"OP":         threadData.OP,
		"Replies":    threadData.Replies,
	})
}

func (h *WebHandler) reply(c *gin.Context) {
	board := c.Param("board")
	rawThread := c.Param("thread")
	thread, err := strconv.Atoi(rawThread)
	if err != nil {
		err := model.NewNotFound("page", rawThread)
		c.JSON(model.Status(err), gin.H{
			"status": model.Status(err),
			"result": err,
		})
		return
	}
	com := c.PostForm("text")
	newPost, err := h.userspace.Reply(c.Request.Context(), board, com, thread)
	if err != nil {
		c.JSON(model.Status(err), gin.H{
			"status": model.Status(err),
			"result": err,
		})
		return
	}
	c.Redirect(http.StatusFound, fmt.Sprintf("/%s/thread/%v#%v", board, thread, newPost.No))
}

func (h *WebHandler) newthread(c *gin.Context) {
	board := c.Param("board")
	com := c.PostForm("text")
	fmt.Printf("Newthread: %s", com)
	newPost, err := h.userspace.NewThread(c.Request.Context(), board, com)
	if err != nil {
		c.JSON(model.Status(err), gin.H{
			"status": model.Status(err),
			"result": err,
		})
		return
	}
	c.Redirect(http.StatusFound, fmt.Sprintf("/%s/thread/%v", board, newPost.No))
}
