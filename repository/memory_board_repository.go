package repository

import (
	"sync"

	"github.com/1ort/goimbo/model"
)

type Board model.Board

type MemoryBoardRepository struct {
	Boards []*Board
	mutex  sync.Mutex
}

func (self *MemoryBoardRepository) NewBoard(slug, name, descr string) (*Board, error) {
	if ex, _ := self.IsBoardExists(slug); ex {
		return nil, model.NewConflict("board", slug)
	}
	b := &Board{
		Slug:  slug,
		Name:  name,
		Descr: descr,
	}
	self.mutex.Lock()
	defer self.mutex.Unlock()
	self.Boards = append(self.Boards, b)
	return b, nil
}

func (self *MemoryBoardRepository) GetBoard(slug string) (*Board, error) {
	self.mutex.Lock()
	defer self.mutex.Unlock()
	for _, b := range self.Boards {
		if b.Slug == slug {
			return b, nil
		}
	}
	return nil, model.NewNotFound("board", slug)
}

func (self *MemoryBoardRepository) IsBoardExists(slug string) (bool, error) {
	self.mutex.Lock()
	defer self.mutex.Unlock()
	for _, b := range self.Boards {
		if b.Slug == slug {
			return true, nil
		}
	}
	return false, nil
}

func (self *MemoryBoardRepository) GetBoardList() ([]*Board, error) {
	self.mutex.Lock()
	defer self.mutex.Unlock()
	return self.Boards, nil
}
