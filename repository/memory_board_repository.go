package repository

import (
	"sync"

	"context"

	"github.com/1ort/goimbo/model"
)

type MemoryBoardRepository struct {
	Boards []*model.Board
	mutex  sync.Mutex
}

func (self *MemoryBoardRepository) NewBoard(ctx context.Context, slug, name, descr string) (*model.Board, error) {
	if ex, _ := self.IsBoardExists(ctx, slug); ex {
		return nil, model.NewConflict("board", slug)
	}
	b := &model.Board{
		Slug:  slug,
		Name:  name,
		Descr: descr,
	}
	self.mutex.Lock()
	defer self.mutex.Unlock()
	self.Boards = append(self.Boards, b)
	return b, nil
}

func (self *MemoryBoardRepository) GetBoard(ctx context.Context, slug string) (*model.Board, error) {
	self.mutex.Lock()
	defer self.mutex.Unlock()
	for _, b := range self.Boards {
		if b.Slug == slug {
			return b, nil
		}
	}
	return nil, model.NewNotFound("board", slug)
}

func (self *MemoryBoardRepository) IsBoardExists(ctx context.Context, slug string) (bool, error) {
	self.mutex.Lock()
	defer self.mutex.Unlock()
	for _, b := range self.Boards {
		if b.Slug == slug {
			return true, nil
		}
	}
	return false, nil
}

func (self *MemoryBoardRepository) GetBoardList(ctx context.Context) ([]*model.Board, error) {
	self.mutex.Lock()
	defer self.mutex.Unlock()
	return self.Boards, nil
}
