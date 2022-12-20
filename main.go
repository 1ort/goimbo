package main

import (
	"flag"

	"github.com/1ort/goimbo/handler"
	"github.com/1ort/goimbo/model"
	"github.com/1ort/goimbo/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	confPtr := flag.String("config", "config.yaml", "config file path")
	flag.Parse()
	config := ReadConfig(*confPtr)

	router := gin.Default()

	init_boards := []*model.Board{
		{
			Slug:  "po",
			Name:  "Politics",
			Descr: "Политика",
		},
		{
			Slug:  "b",
			Name:  "bред",
			Descr: "Бредач",
		},
		{
			Slug:  "r",
			Name:  "Random",
			Descr: "Рандомач",
		},
		{
			Slug:  "vg",
			Name:  "Video games general",
			Descr: "Майнкрафт и дота",
		},
	}

	boardRepo := repository.NewMemoryBoardRepository(
		&repository.MemoryBoardRepositoryConfig{
			Boards: init_boards,
		},
	)
	postRepo := repository.NewMemoryPostRepository(
		&repository.MemoryPostRepositoryConfig{
			BoardRepo: boardRepo,
			Posts:     make(map[string][]*model.Post),
		},
	)

	handlerConfig := handler.HandlerConfig{
		R:         router,
		BaseUrl:   config.GetBaseApiUrl(),
		BoardRepo: boardRepo,
		PostRepo:  postRepo,
	}

	handler.NewHandler(&handlerConfig)

	router.Run(config.GetAppAddr())

	// db.InitDatabase(db_pool)

	//framework.SetupMiddlewares(app)
	//app.Use(framework.DBConnPool(db_pool))

	//handler.ApplyHandlers(app)
	//framework.RunServer(app, config.GetAppAddr())
}
