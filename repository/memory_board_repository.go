package repository

import (
	"sync"

	"context"

	"github.com/1ort/goimbo/model"
)

type memoryBoardRepository struct {
	Boards []*model.Board
	mutex  sync.Mutex
}

type MemoryBoardRepositoryConfig struct {
	Boards []*model.Board
}

func NewMemoryBoardRepository(c *MemoryBoardRepositoryConfig) model.BoardRepository {
	return &memoryBoardRepository{
		Boards: c.Boards,
		mutex:  sync.Mutex{},
	}
}

func (mbr *memoryBoardRepository) NewBoard(ctx context.Context, slug, name, descr string) (*model.Board, error) {
	if ok, _ := mbr.GetBoard(ctx, slug); ok != nil {
		return nil, model.NewConflict("board", slug)
	}
	b := &model.Board{
		Slug:  slug,
		Name:  name,
		Descr: descr,
	}
	mbr.mutex.Lock()
	defer mbr.mutex.Unlock()
	mbr.Boards = append(mbr.Boards, b)
	return b, nil
}

func (mbr *memoryBoardRepository) GetBoard(ctx context.Context, slug string) (*model.Board, error) {
	mbr.mutex.Lock()
	defer mbr.mutex.Unlock()
	for _, b := range mbr.Boards {
		if b.Slug == slug {
			return b, nil
		}
	}
	return nil, model.NewNotFound("board", slug)
}

func (mbr *memoryBoardRepository) GetBoardList(ctx context.Context) ([]*model.Board, error) {
	mbr.mutex.Lock()
	defer mbr.mutex.Unlock()
	return mbr.Boards, nil
}
