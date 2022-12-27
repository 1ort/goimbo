package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/1ort/goimbo/model"
	"github.com/gin-gonic/gin"
)

func (h *APIHandler) getBoards(c *gin.Context) {
	boardList, err := h.userspace.GetBoards(c.Request.Context())
	if err != nil {
		c.JSON(model.Status(err), gin.H{
			"status": model.Status(err),
			"result": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"result": boardList,
	})
}

func (h *APIHandler) getBoard(c *gin.Context) {
	board := c.Param("board")
	boardData, err := h.userspace.GetBoard(c.Request.Context(), board)
	if err != nil {
		c.JSON(model.Status(err), gin.H{
			"status": model.Status(err),
			"result": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"result": boardData,
	})
}

func (h *APIHandler) getPage(c *gin.Context) {
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
	_, err = h.userspace.GetBoard(c.Request.Context(), board) //404 if board does not exist
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
	{
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"result": pageData,
		})
	}
}

func (h *APIHandler) getThread(c *gin.Context) {
	board := c.Param("board")
	rawThread := c.Param("op")
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
	{
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"result": threadData,
		})
	}
}
