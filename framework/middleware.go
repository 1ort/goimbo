package framework

import (
	"github.com/gin-gonic/gin"
)

func SetupMiddlewares(app *gin.Engine) {
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
}
