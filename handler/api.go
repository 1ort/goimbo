package handler

import (
	"net/http"
	"strconv"

	"github.com/1ort/goimbo/model"
	"github.com/gin-gonic/gin"
)

type APIConfig struct {
	R         *gin.Engine //router
	BaseURL   string
	Userspace model.Userspace
}

type APIHandler struct {
	userspace model.Userspace
	r         *gin.Engine
}

func SetAPIHandler(cfg *APIConfig) {
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

func (h *APIHandler) handleError(c *gin.Context, err error) bool {
	if err == nil {
		return false
	}
	m := gin.H{
		"status":  model.Status(err),
		"message": err,
	}
	if model.Status(err) == 500 {
		m["message"] = "Internal error"
	}
	c.JSON(model.Status(err), m)
	return true
}

func (h *APIHandler) getBoards(c *gin.Context) {
	boardList, err := h.userspace.GetBoards(c.Request.Context())
	if handled := h.handleError(c, err); handled {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"result": boardList,
	})
}

func (h *APIHandler) getBoard(c *gin.Context) {
	var br BoardRequest
	if err := c.ShouldBindUri(&br); err != nil {
		h.handleError(c, model.NewBadRequest("Invalid board"))
		return
	}
	boardData, err := h.userspace.GetBoard(c.Request.Context(), br.Board)
	if handled := h.handleError(c, err); handled {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"result": boardData,
	})
}

func (h *APIHandler) getPage(c *gin.Context) {
	var bp BoardPageRequest
	if err := c.ShouldBindUri(&bp); err != nil {
		h.handleError(c, model.NewBadRequest("Invalid board or page"))
		return
	}
	_, err := h.userspace.GetBoard(c.Request.Context(), bp.Board) //404 if board does not exist
	if handled := h.handleError(c, err); handled {
		return
	}
	pageData, err := h.userspace.GetBoardPage(c.Request.Context(), bp.Board, bp.Page)
	if handled := h.handleError(c, err); handled {
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
	var t ThreadPageRequest
	if err := c.ShouldBindUri(&t); err != nil {
		h.handleError(c, model.NewBadRequest("Invalid board or thread number"))
		return
	}
	threadData, err := h.userspace.GetThread(c.Request.Context(), t.Board, t.Thread)
	if err != nil {
		err := model.NewNotFound("thread", strconv.Itoa(t.Thread))
		h.handleError(c, err)
		return
	}
	{
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"result": threadData,
		})
	}
}
