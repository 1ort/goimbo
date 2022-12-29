package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/1ort/goimbo/handler"
	"github.com/1ort/goimbo/model"
	"github.com/1ort/goimbo/repository"
	"github.com/1ort/goimbo/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	confPtr := flag.String("config", "config.yaml", "config file path")
	flag.Parse()

	config := ReadConfig(*confPtr)

	pgpool, err := pgxpool.New(context.Background(), config.GetDataBaseURL())
	if err != nil {
		panic(err)
	}

	boardrepo := repository.NewMemoryBoardRepository(
		&repository.MemoryBoardRepositoryConfig{
			Boards: []*model.Board{
				{Slug: "b", Name: "Board B", Descr: "Board B description"},
				{Slug: "a", Name: "Board A", Descr: "Board A description"},
				{Slug: "c", Name: "Board C", Descr: "Board C description"},
			},
		})
	postrepo := repository.NewPGPostRepository(
		&repository.PgPostRepoConfig{
			Pool: pgpool,
		})

	userspace := service.NewUserspaceService(
		&service.UserspaceServiceConfig{
			BoardRepository: boardrepo,
			PostRepository:  postrepo,
		})

	if !config.Web.Enabled && !config.API.Enabled {
		fmt.Println("At least one: Web or API must be enabled. Enable in config")
		return
	}
	router := gin.Default()
	if config.API.Enabled {
		handler.SetAPIHandler(
			&handler.APIConfig{
				R:         router,
				BaseURL:   config.API.BaseURL,
				Userspace: userspace,
			})
	}
	if config.Web.Enabled {
		handler.SetWebHandler(
			&handler.WebConfig{
				R:             router,
				BaseURL:       config.Web.BaseURL,
				Userspace:     userspace,
				CookieSecret:  config.Web.CookieSecret,
				XCSRFSecret:   config.Web.XCSRFSecret,
				EnableCaptcha: true,
			})
	}
	router.Run(config.GetAppAddr()) //nolint
}
