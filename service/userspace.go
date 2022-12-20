package service

import (
	"context"

	"github.com/1ort/goimbo/model"
)

type UserspaceService struct {
	PostRepository  model.PostRepository
	BoardRepository model.BoardRepository
}

type UserspaceServiceConfig struct {
	PostRepository  model.PostRepository
	BoardRepository model.BoardRepository
}

func NewUserspaceService(c *UserspaceServiceConfig) model.Userspace {
	return &UserspaceService{
		PostRepository:  c.PostRepository,
		BoardRepository: c.BoardRepository,
	}
}

func (self *UserspaceService) Boards(ctx context.Context) ([]*model.Board, error) {
	return self.BoardRepository.GetBoardList(ctx)
	//return nil, nil
}
func (self *UserspaceService) Threads(ctx context.Context, board string) ([]*model.ThreadListPage, error) {
	return nil, nil
}
func (self *UserspaceService) Catalog(ctx context.Context, board string) ([]*model.CatalogPage, error) {
	return nil, nil
}
func (self *UserspaceService) Index(ctx context.Context, board string, page int) (*model.ThreadPage, error) {
	return nil, nil
}
func (self *UserspaceService) Thread(ctx context.Context, board string, op int) ([]*model.Post, error) {
	return nil, nil
}
func (self *UserspaceService) NewPost(ctx context.Context, board string, resto int, com string) (*model.Post, error) {
	return nil, nil
}
