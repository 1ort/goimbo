package handler

import (
	"fmt"
	"net/http"

	"github.com/1ort/goimbo/db"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewBoard(c *gin.Context) {

	db_pool := c.MustGet("db").(*pgxpool.Pool)
	slug := c.Query("slug")
	name := c.DefaultQuery("name", slug)
	descr := c.DefaultQuery("descr", "New board")

	err := db.NewBoard(db_pool, slug, name, descr)
	if err != nil {
		fmt.Printf("%e", err)
		c.String(http.StatusInternalServerError, "Can not create new board %s", slug)
	} else {
		c.String(http.StatusOK, "Successfully created new board %s", slug)
	}
}
