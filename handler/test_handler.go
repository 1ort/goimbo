package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func get_boards(c *gin.Context) {
	c.String(http.StatusOK, "Board list here")
}

func get_threads(c *gin.Context) {
	board := c.Param("board")
	c.String(http.StatusOK, "Thread list on %s board here", board)

}

func get_catalog(c *gin.Context) {
	board := c.Param("board")
	c.String(http.StatusOK, "Catalog of %s board here", board)

}

func get_archive(c *gin.Context) {
	board := c.Param("board")
	c.String(http.StatusOK, "Archive of %s board here", board)
}

func get_page(c *gin.Context) {
	board := c.Param("board")
	page := c.Param("page")
	if _, err := strconv.Atoi(page); err != nil {
		c.String(http.StatusBadRequest, "Page number must be an integer, got %s instead", page)
	} else {
		c.String(http.StatusOK, "Page number %s on %s board here", page, board)
	}
}

func get_thread(c *gin.Context) {
	board := c.Param("board")
	op := c.Param("op")
	if _, err := strconv.Atoi(op); err != nil {
		c.String(http.StatusBadRequest, "Thread number must be an integer, got %s instead", op)
	} else {
		c.String(http.StatusOK, "Thread number %s content on %s board here", op, board)
	}
}

func ApplyHandlers(app *gin.Engine) {
	//Default has logger and recovery (crash-free) middleware
	//router := gin.Default()

	app.GET("/boards", get_boards)
	app.GET("/:board/threads", get_threads)
	app.GET("/:board/catalog", get_catalog)
	app.GET("/:board/archive", get_archive)
	app.GET("/:board/:page", get_page)
	app.GET("/:board/thread/:op", get_thread)
	app.POST("/:board/newpost", NewPost)
	app.GET("/newboard", NewBoard)
	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	//router.Run()
}
