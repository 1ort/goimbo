package handler

import (
	"fmt"
	"net/http"

	"github.com/1ort/goimbo/model"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

// TODO: Put it all in separate files and add the bindall function
type CaptchaRequest struct {
	ID       string `form:"captchaId" binding:"required"`
	Solution string `form:"captchaSolution" binding:"required"`
}

type PostRequest struct {
	Com string `form:"text" binding:"required"`
}

type BoardRequest struct {
	Board string `uri:"board" binding:"required"`
}

type BoardPageRequest struct {
	BoardRequest
	Page int `uri:"page" binding:"gte=0"`
}

type ThreadPageRequest struct {
	BoardRequest
	Thread int `uri:"thread" binding:"gte=0"`
}

func (h *WebHandler) handleError(c *gin.Context, err error) bool {
	if err == nil {
		return false
	}
	m := gin.H{
		"status":  model.Status(err),
		"message": err,
	}
	if model.Status(err) == 500 {
		m["message"] = "Something went wrong"
	}
	c.HTML(model.Status(err), "error.page.tmpl", m)
	return true
}

// TODO: Move captcha to a separate interface
func (h *WebHandler) verifyCaptcha(c *gin.Context) error {
	if !h.enableCaptcha {
		return nil
	}
	var cr CaptchaRequest
	if err := c.ShouldBind(&cr); err != nil {
		return model.NewBadRequest("Captcha error")
	}
	verified := captcha.VerifyString(cr.ID, cr.Solution)
	if !verified {
		return model.NewBadRequest("Incorrect captcha")
	}
	return nil
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

// TODO: Do not create captcha if disabled
func (h *WebHandler) boardPage(c *gin.Context) {
	var p BoardPageRequest
	if err := c.ShouldBindUri(&p); err != nil {
		h.handleError(c, model.NewBadRequest("Invalid board or page"))
		return
	}
	boardData, err := h.userspace.GetBoard(c.Request.Context(), p.Board)
	if handled := h.handleError(c, err); handled {
		return
	}
	pageData, err := h.userspace.GetBoardPage(c.Request.Context(), p.Board, p.Page)
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
	var p ThreadPageRequest
	if err := c.ShouldBindUri(&p); err != nil {
		h.handleError(c, model.NewBadRequest("Invalid thread"))
		return
	}
	boardData, err := h.userspace.GetBoard(c.Request.Context(), p.Board)
	if handled := h.handleError(c, err); handled {
		return
	}
	threadData, err := h.userspace.GetThread(c.Request.Context(), p.Board, p.Thread)
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
	var t ThreadPageRequest
	var p PostRequest

	if err := c.ShouldBindUri(&t); err != nil {
		h.handleError(c, model.NewBadRequest("Invalid thread"))
		return
	}
	if err := c.ShouldBind(&p); err != nil {
		h.handleError(c, model.NewBadRequest("Invalid post content"))
		return
	}

	err := h.verifyCaptcha(c)
	if handled := h.handleError(c, err); handled {
		return
	}

	newPost, err := h.userspace.Reply(c.Request.Context(), t.Board, p.Com, t.Thread)
	if handled := h.handleError(c, err); handled {
		return
	}
	c.Redirect(http.StatusFound, fmt.Sprintf("/%s/thread/%v#%v", t.Board, t.Thread, newPost.No))
}

func (h *WebHandler) newthread(c *gin.Context) {
	var b BoardRequest
	var p PostRequest

	if err := c.ShouldBindUri(&b); err != nil {
		h.handleError(c, model.NewBadRequest("Invalid board"))
		return
	}
	if err := c.ShouldBind(&p); err != nil {
		h.handleError(c, model.NewBadRequest("Invalid post content"))
		return
	}

	err := h.verifyCaptcha(c)
	if handled := h.handleError(c, err); handled {
		return
	}

	newPost, err := h.userspace.NewThread(c.Request.Context(), b.Board, p.Com)
	if handled := h.handleError(c, err); handled {
		return
	}
	c.Redirect(http.StatusFound, fmt.Sprintf("/%s/thread/%v", b.Board, newPost.No))
}
