package main

import (
	"github.com/1ort/goimbo/db"
	"github.com/1ort/goimbo/framework"
	"github.com/1ort/goimbo/handler"
)

func main() {
	app := framework.NewApp()

	db_pool := db.NewPool("postgres://postgres:postgres@localhost:5432/goimbo") //TODO: use config file / env vars
	defer db_pool.Close()

	db.InitDatabase(db_pool)

	framework.SetupMiddlewares(app)
	app.Use(framework.DBConnPool(db_pool))

	handler.ApplyHandlers(app)
	framework.RunServer(app, ":3000")
}
