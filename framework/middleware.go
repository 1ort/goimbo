package framework

import (
	"net/http"

	"github.com/1ort/goimbo/db"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func DBConnPool(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", pool)
		c.Next()
	}
}

func BoardExists() gin.HandlerFunc {
	return func(c *gin.Context) {
		db_pool := c.MustGet("db").(*pgxpool.Pool)
		board := c.Param("board")
		res, err := db.BoardExists(db_pool, board)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		} else if !res {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.Next()
		}

	}
}

func SetupMiddlewares(app *gin.Engine) {
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
}
