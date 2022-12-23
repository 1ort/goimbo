package main

import (
	"flag"
	"fmt"

	"github.com/1ort/goimbo/handler"
	"github.com/1ort/goimbo/service"
	"github.com/gin-gonic/gin"
)

func main() {
	confPtr := flag.String("config", "config.yaml", "config file path")
	flag.Parse()

	config := ReadConfig(*confPtr)

	// boardRepo := repository.NewMemoryBoardRepository(
	// 	&repository.MemoryBoardRepositoryConfig{
	// 		Boards: init_boards,
	// 	})

	// postRepo := repository.NewMemoryPostRepository(
	// 	&repository.MemoryPostRepositoryConfig{
	// 		Posts: make(map[string][]*model.Post),
	// 	})

	// userspace := service.NewUserspaceService(
	// 	&service.UserspaceServiceConfig{
	// 		PostRepository:  postRepo,
	// 		BoardRepository: boardRepo,
	// 	})

	userspace := service.NewMockUserspace()

	if !config.Web.Enabled && !config.Api.Enabled {
		fmt.Println("At least one: Web or API must be enabled. Enable in config")
		return
	}
	router := gin.Default()
	if config.Web.Enabled {
		handler.SetWebHandler(
			&handler.HandlerConfig{
				R:         router,
				BaseUrl:   config.Web.BaseUrl,
				Userspace: userspace,
			})
	}
	if config.Api.Enabled {
		handler.SetApiHandler(
			&handler.HandlerConfig{
				R:         router,
				BaseUrl:   config.Api.BaseUrl,
				Userspace: userspace,
			})
	}

	router.Run(config.GetAppAddr())
}
