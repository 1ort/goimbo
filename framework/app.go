package framework

import (
	"github.com/gin-gonic/gin"
)

func NewApp() *gin.Engine {
	app := gin.New()
	return app
}

func RunServer(app *gin.Engine, port string) {
	app.Run(port)
}
