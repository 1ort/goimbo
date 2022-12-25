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

func (u *UserspaceService) GetBoard(ctx context.Context, slug string) (*model.Board, error) {
	return nil, nil
}
func (u *UserspaceService) GetBoards(ctx context.Context) ([]*model.Board, error) {
	return u.BoardRepository.GetBoardList(ctx)
}
func (u *UserspaceService) GetThread(ctx context.Context, board string, no int) (*model.Thread, error) {
	return nil, nil
}
func (u *UserspaceService) GetThreadPreview(ctx context.Context, board string, no int) (*model.ThreadPreview, error) {
	return nil, nil
}
func (u *UserspaceService) GetBoardPage(ctx context.Context, board string, page int) (*model.BoardPage, error) {
	return nil, nil
}
func (u *UserspaceService) NewThread(ctx context.Context, board, com string) error {
	return nil
}
func (u *UserspaceService) Reply(ctx context.Context, board, com string, parent int) error {
	return nil
}
