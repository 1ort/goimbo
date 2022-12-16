package main

import (
	"github.com/1ort/goimbo/db"
	"github.com/1ort/goimbo/framework"
	"github.com/1ort/goimbo/handler"
	"github.com/1ort/goimbo/utils"
)

func main() {
	app := framework.NewApp()
	db_url := utils.GetDataBaseUrl()
	db_pool := db.NewPool(db_url)
	defer db_pool.Close()

	db.InitDatabase(db_pool)

	framework.SetupMiddlewares(app)
	app.Use(framework.DBConnPool(db_pool))

	handler.ApplyHandlers(app)
	framework.RunServer(app, ":3000")
}
