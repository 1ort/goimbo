package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/1ort/goimbo/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) get_boards(c *gin.Context) {
	ctx := c.Request.Context()
	boardList, err := h.userspace.Boards(ctx)
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

func (h *Handler) get_threads(c *gin.Context) {
	ctx := c.Request.Context()
	board := c.Param("board")
	threadList, err := h.userspace.Threads(ctx, board)
	if err != nil {
		c.JSON(model.Status(err), gin.H{
			"status": model.Status(err),
			"result": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"result": threadList,
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
