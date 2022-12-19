package repository

import (
	"fmt"
	"time"

	"github.com/1ort/goimbo/model"
)

type Post model.Post

type MemoryPostRepository struct {
	BoardRepo model.BoardRepository
	Posts     map[string][]*Post
}

func (self *MemoryPostRepository) NewPost(resto int, board, com string) (*Post, error) {
	if _, err := self.BoardRepo.GetBoard(board); err != nil {
		return nil, err
	}
	if resto != 0 {
		if _, err := self.GetPost(resto, board); err != nil {
			return nil, err
		}
	}
	p := &Post{
		No:    len(self.Posts[board]) + 1,
		Resto: resto,
		Board: board,
		Com:   com,
		Time:  int(time.Now().Unix()),
	}
	self.Posts[board] = append(self.Posts[board], p)
	return p, nil
}

func (self *MemoryPostRepository) GetPost(no int, board string) (*Post, error) {
	for _, p := range self.Posts[board] {
		if p.No == no {
			return p, nil
		}
	}
	return nil, model.NewNotFound("post", fmt.Sprintf("%d", no))
}

func (self *MemoryPostRepository) GetThreadHistory(no int, board string) ([]*Post, error) {
	var posts []*Post
	for _, p := range self.Posts[board] {
		if p.No == no || p.Resto == no {
			posts = append(posts, p)
		}
	}
	return posts, nil
}

func (self *MemoryPostRepository) DeletePost(no int, board string) (bool, error) {
	for i, p := range self.Posts[board] {
		if p.No == no {
			self.Posts[board] = append(self.Posts[board][:i], self.Posts[board][i+1:]...)
			return true, nil
		}
	}
	return false, model.NewNotFound("post", fmt.Sprintf("%d", no))
}

func (self *MemoryPostRepository) IsOp(no int, board string) (bool, error) {
	for _, p := range self.Posts[board] {
		if p.No == no {
			return p.Resto == 0, nil
		}
	}
	return false, model.NewNotFound("post", fmt.Sprintf("%d", no))
}
