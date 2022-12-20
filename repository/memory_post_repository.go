package repository

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/1ort/goimbo/model"
)

type memoryPostRepository struct {
	BoardRepo model.BoardRepository
	Posts     map[string][]*model.Post
	mutex     sync.Mutex
}

type MemoryPostRepositoryConfig struct {
	BoardRepo model.BoardRepository
	Posts     map[string][]*model.Post
}

func NewMemoryPostRepository(c *MemoryPostRepositoryConfig) model.PostRepository {
	return &memoryPostRepository{
		BoardRepo: c.BoardRepo,
		Posts:     c.Posts,
		mutex:     sync.Mutex{},
	}
}

func (self *memoryPostRepository) GetThreadList(ctx context.Context, board string) ([]*model.Post, error) {
	if _, err := self.BoardRepo.GetBoard(ctx, board); err != nil {
		return nil, err
	}
	self.mutex.Lock()
	defer self.mutex.Unlock()
	var posts []*model.Post
	for _, p := range self.Posts[board] {
		if p.Resto == 0 {
			posts = append(posts, p)
		}
	}
	return posts, nil
}

func (self *memoryPostRepository) NewPost(ctx context.Context, resto int, board, com string) (*model.Post, error) {
	if _, err := self.BoardRepo.GetBoard(ctx, board); err != nil {
		return nil, err
	}
	if resto != 0 {
		isop, err := self.IsOp(ctx, resto, board)
		if err != nil {
			return nil, err
		}
		if !isop {
			return nil, model.NewBadRequest(fmt.Sprintf("post %v on board %v is not OP", resto, board))
		}
	}
	self.mutex.Lock()
	defer self.mutex.Unlock()
	p := &model.Post{
		No:    len(self.Posts[board]) + 1,
		Resto: resto,
		Board: board,
		Com:   com,
		Time:  int(time.Now().Unix()),
	}
	self.Posts[board] = append(self.Posts[board], p)
	return p, nil
}

func (self *memoryPostRepository) GetPost(ctx context.Context, no int, board string) (*model.Post, error) {
	if _, err := self.BoardRepo.GetBoard(ctx, board); err != nil {
		return nil, err
	}
	self.mutex.Lock()
	defer self.mutex.Unlock()
	for _, p := range self.Posts[board] {
		if p.No == no {
			return p, nil
		}
	}
	return nil, model.NewNotFound("post", fmt.Sprintf("%d", no))
}

func (self *memoryPostRepository) GetThreadHistory(ctx context.Context, no int, board string) ([]*model.Post, error) {
	if _, err := self.BoardRepo.GetBoard(ctx, board); err != nil {
		return nil, err
	}

	isop, err := self.IsOp(ctx, no, board)
	if err != nil {
		return nil, err
	}

	if !isop {
		return nil, model.NewBadRequest(fmt.Sprintf("post %v on board %v is not OP", no, board))
	}

	self.mutex.Lock()
	defer self.mutex.Unlock()
	var posts []*model.Post
	for _, p := range self.Posts[board] {
		if p.No == no || p.Resto == no {
			posts = append(posts, p)
		}
	}
	return posts, nil
}

func (self *memoryPostRepository) DeletePost(ctx context.Context, no int, board string) (bool, error) {
	if _, err := self.BoardRepo.GetBoard(ctx, board); err != nil {
		return false, err
	}
	self.mutex.Lock()
	defer self.mutex.Unlock()
	for _, p := range self.Posts[board] {
		if p.No == no {
			p.Com = "content deleted"
			return true, nil
		}
	}
	return false, model.NewNotFound("post", fmt.Sprintf("%d", no))
}

func (self *memoryPostRepository) IsOp(ctx context.Context, no int, board string) (bool, error) {
	if _, err := self.BoardRepo.GetBoard(ctx, board); err != nil {
		return false, err
	}
	self.mutex.Lock()
	defer self.mutex.Unlock()
	for _, p := range self.Posts[board] {
		if p.No == no {
			return p.Resto == 0, nil
		}
	}
	return false, model.NewNotFound("post", fmt.Sprintf("%d", no))
}
