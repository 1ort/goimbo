package handler

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"

	"github.com/1ort/goimbo/db"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPost(c *gin.Context) {

	db_pool := c.MustGet("db").(*pgxpool.Pool)
	board := c.Param("board")
	resto_str := c.DefaultQuery("resto", "0")

	resto, err := strconv.Atoi(resto_str)
	if err != nil {
		c.String(http.StatusBadRequest, "Post number must be digit, got %s", resto_str)
		return
	}

	if !db.IsOp(db_pool, board, resto) {
		c.String(http.StatusBadRequest, "Post #%s is not OP or the thread doesn't exist", resto_str)
		return
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	com := buf.String()

	err = db.InsertPost(db_pool, resto, board, com)
	if err != nil {
		fmt.Printf("%e", err)
		c.String(http.StatusInternalServerError, "Can not post")
	} else {
		c.String(http.StatusOK, "Successfully posted!")
	}
}
