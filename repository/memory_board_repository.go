package repository

import "github.com/1ort/goimbo/model"

type Board model.Board

type MemoryBoardRepository struct {
	Boards []*Board
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
	self.Boards = append(self.Boards, b)
	return b, nil
}

func (self *MemoryBoardRepository) GetBoard(slug string) (*Board, error) {
	for _, b := range self.Boards {
		if b.Slug == slug {
			return b, nil
		}
	}
	return nil, model.NewNotFound("board", slug)
}

func (self *MemoryBoardRepository) IsBoardExists(slug string) (bool, error) {
	for _, b := range self.Boards {
		if b.Slug == slug {
			return true, nil
		}
	}
	return false, nil
}

func (self *MemoryBoardRepository) GetBoardList() ([]*Board, error) {
	return self.Boards, nil
}
