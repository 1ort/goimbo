package main

import (
	"flag"

	"github.com/1ort/goimbo/handler"
	"github.com/1ort/goimbo/model"
	"github.com/1ort/goimbo/repository"
	"github.com/1ort/goimbo/service"
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
		})

	postRepo := repository.NewMemoryPostRepository(
		&repository.MemoryPostRepositoryConfig{
			Posts: make(map[string][]*model.Post),
		})

	userspace := service.NewUserspaceService(
		&service.UserspaceServiceConfig{
			PostRepository:  postRepo,
			BoardRepository: boardRepo,
		})

	handler.NewHandler(
		&handler.HandlerConfig{
			R:          router,
			ApiBaseUrl: config.GetBaseApiUrl(),
			WebBaseUrl: config.GetBaseWebUrl(),
			Userspace:  userspace,
		})

	router.Run(config.GetAppAddr())
}
