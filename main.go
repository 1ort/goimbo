package main

import (
	"github.com/1ort/goimbo/framework"
	"github.com/1ort/goimbo/handler"
)

func main() {
	app := framework.NewApp()
	framework.SetupMiddlewares(app)
	handler.ApplyHandlers(app)
	framework.RunServer(app, ":3000")
}
