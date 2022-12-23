package repository

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/1ort/goimbo/model"
)

type memoryPostRepository struct {
	Posts map[string][]*model.Post
	mutex sync.Mutex
}

type MemoryPostRepositoryConfig struct {
	Posts map[string][]*model.Post
}

func NewMemoryPostRepository(c *MemoryPostRepositoryConfig) model.PostRepository {
	return &memoryPostRepository{
		Posts: c.Posts,
		mutex: sync.Mutex{},
	}
}

func (self *memoryPostRepository) GetThreadList(ctx context.Context, board string) ([]*model.Post, error) {
	self.mutex.Lock()
	defer self.mutex.Unlock()
	var posts []*model.Post
	for _, p := range self.Posts[board] {
		if p.Parent == 0 {
			posts = append(posts, p)
		}
	}
	return posts, nil
}

func (self *memoryPostRepository) NewPost(ctx context.Context, parent int, board, com string) (*model.Post, error) {
	if parent != 0 {
		isop, err := self.IsOp(ctx, parent, board)
		if err != nil {
			return nil, err
		}
		if !isop {
			return nil, model.NewBadRequest(fmt.Sprintf("post %v on board %v is not OP", parent, board))
		}
	}
	self.mutex.Lock()
	defer self.mutex.Unlock()
	p := &model.Post{
		No:     len(self.Posts[board]) + 1,
		Parent: parent,
		Board:  board,
		Com:    com,
		Time:   int(time.Now().Unix()),
	}
	self.Posts[board] = append(self.Posts[board], p)
	return p, nil
}

func (self *memoryPostRepository) GetPost(ctx context.Context, no int, board string) (*model.Post, error) {
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
		if p.No == no || p.Parent == no {
			posts = append(posts, p)
		}
	}
	return posts, nil
}

func (self *memoryPostRepository) DeletePost(ctx context.Context, no int, board string) (bool, error) {
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
	self.mutex.Lock()
	defer self.mutex.Unlock()
	for _, p := range self.Posts[board] {
		if p.No == no {
			return p.Parent == 0, nil
		}
	}
	return false, model.NewNotFound("post", fmt.Sprintf("%d", no))
}
