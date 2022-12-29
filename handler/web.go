package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/1ort/goimbo/model"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

func (h *WebHandler) handleError(c *gin.Context, err error) bool {
	if err == nil {
		return false
	}
	m := gin.H{
		"status":  model.Status(err),
		"message": err,
	}
	if model.Status(err) == 500 {
		m["status"] = "Something went wrong"
	}
	c.HTML(model.Status(err), "error.page.tmpl", m)
	return true
}

func (h *WebHandler) mainPage(c *gin.Context) {
	boards, err := h.userspace.GetBoards(c.Request.Context())
	if handled := h.handleError(c, err); handled {
		return
	}
	c.HTML(http.StatusOK, "dir.page.tmpl", gin.H{
		"Boards":      boards,
		"XCSRF_TOKEN": csrf.GetToken(c),
	})
}

func (h *WebHandler) redirectToZeroPage(c *gin.Context) {
	board := c.Param("board")
	c.Redirect(http.StatusFound, fmt.Sprintf("/%s/page/0", board))
}

func (h *WebHandler) boardPage(c *gin.Context) {
	board := c.Param("board")
	rawPage := c.Param("page")
	page, err := strconv.Atoi(rawPage)
	if err != nil {
		err := model.NewNotFound("page", rawPage)
		h.handleError(c, err)
		return
	}
	boardData, err := h.userspace.GetBoard(c.Request.Context(), board)
	if handled := h.handleError(c, err); handled {
		return
	}
	pageData, err := h.userspace.GetBoardPage(c.Request.Context(), board, page)
	if handled := h.handleError(c, err); handled {
		return
	}
	c.HTML(http.StatusOK, "board.page.tmpl", gin.H{
		"Page":        pageData.Page,
		"Threads":     pageData.Threads,
		"Board":       boardData,
		"XCSRF_TOKEN": csrf.GetToken(c),
		"captcha": gin.H{
			"ID":      captcha.New(),
			"enabled": h.enableCaptcha,
		},
	})
}

func (h *WebHandler) threadPage(c *gin.Context) {
	board := c.Param("board")
	rawThread := c.Param("thread")
	boardData, err := h.userspace.GetBoard(c.Request.Context(), board)
	if handled := h.handleError(c, err); handled {
		return
	}
	thread, err := strconv.Atoi(rawThread)
	if err != nil {
		err := model.NewNotFound("thread", rawThread)
		h.handleError(c, err)
		return
	}
	threadData, err := h.userspace.GetThread(c.Request.Context(), board, thread)
	if handled := h.handleError(c, err); handled {
		return
	}
	xCSRFToken := csrf.GetToken(c)
	c.HTML(http.StatusOK, "thread.page.tmpl", gin.H{
		"board_data":  boardData,
		"OP":          threadData.OP,
		"Replies":     threadData.Replies,
		"XCSRF_TOKEN": xCSRFToken,
		"captcha": gin.H{
			"ID":      captcha.New(),
			"enabled": h.enableCaptcha,
		},
	})
}

func (h *WebHandler) reply(c *gin.Context) {
	board := c.Param("board")
	rawThread := c.Param("thread")
	thread, err := strconv.Atoi(rawThread)
	if err != nil {
		err := model.NewNotFound("thread", rawThread)
		h.handleError(c, err)
		return
	}
	captchaID := c.PostForm("captchaId")
	captchaSolution := c.PostForm("captchaSolution")
	verified := captcha.VerifyString(captchaID, captchaSolution)
	if !verified {
		err := model.NewBadRequest("Incorrect captcha")
		h.handleError(c, err)
		return
	}
	com := c.PostForm("text")
	newPost, err := h.userspace.Reply(c.Request.Context(), board, com, thread)
	if handled := h.handleError(c, err); handled {
		return
	}
	c.Redirect(http.StatusFound, fmt.Sprintf("/%s/thread/%v#%v", board, thread, newPost.No))
}

func (h *WebHandler) newthread(c *gin.Context) {
	board := c.Param("board")
	com := c.PostForm("text")
	captchaID := c.PostForm("captchaId")
	captchaSolution := c.PostForm("captchaSolution")
	verified := captcha.VerifyString(captchaID, captchaSolution)
	if !verified {
		err := model.NewBadRequest("Captcha incorrect")
		h.handleError(c, err)
		return
	}
	newPost, err := h.userspace.NewThread(c.Request.Context(), board, com)
	if handled := h.handleError(c, err); handled {
		return
	}
	c.Redirect(http.StatusFound, fmt.Sprintf("/%s/thread/%v", board, newPost.No))
}
