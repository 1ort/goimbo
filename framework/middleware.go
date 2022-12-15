package framework

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func DBConnPool(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", pool)
		c.Next()
	}
}

func SetupMiddlewares(app *gin.Engine) {
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
}
