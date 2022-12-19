package main

import (
	"flag"

	//"github.com/1ort/goimbo/db"
	"github.com/1ort/goimbo/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	confPtr := flag.String("config", "config.yaml", "config file path")
	flag.Parse()

	config := ReadConfig(*confPtr)
	// db_url := config.GetDataBaseUrl()

	//app := framework.NewApp()
	// db_pool := db.NewPool(db_url)
	// defer db_pool.Close()

	router := gin.Default()
	handler.NewHandler(&handler.HandlerConfig{
		R:       router,
		BaseUrl: config.GetBaseApiUrl(),
	})

	router.Run(config.GetAppAddr())

	// db.InitDatabase(db_pool)

	//framework.SetupMiddlewares(app)
	//app.Use(framework.DBConnPool(db_pool))

	//handler.ApplyHandlers(app)
	//framework.RunServer(app, config.GetAppAddr())
}
