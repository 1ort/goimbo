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

func (mpr *memoryPostRepository) GetThreadList(ctx context.Context, board string) ([]*model.Post, error) {
	mpr.mutex.Lock()
	defer mpr.mutex.Unlock()
	var posts []*model.Post
	for _, p := range mpr.Posts[board] {
		if p.Parent == 0 {
			posts = append(posts, p)
		}
	}
	return posts, nil
}

func (mpr *memoryPostRepository) NewPost(ctx context.Context, parent int, board, com string) (*model.Post, error) {
	if parent != 0 {
		isop, err := mpr.IsOp(ctx, parent, board)
		if err != nil {
			return nil, err
		}
		if !isop {
			return nil, model.NewBadRequest(fmt.Sprintf("post %v on board %v is not OP", parent, board))
		}
	}
	mpr.mutex.Lock()
	defer mpr.mutex.Unlock()
	p := &model.Post{
		No:     len(mpr.Posts[board]) + 1,
		Parent: parent,
		Board:  board,
		Com:    com,
		Time:   time.Now(),
	}
	mpr.Posts[board] = append(mpr.Posts[board], p)
	return p, nil
}

func (mpr *memoryPostRepository) GetPost(ctx context.Context, no int, board string) (*model.Post, error) {
	mpr.mutex.Lock()
	defer mpr.mutex.Unlock()
	for _, p := range mpr.Posts[board] {
		if p.No == no {
			return p, nil
		}
	}
	return nil, model.NewNotFound("post", fmt.Sprintf("%d", no))
}

func (mpr *memoryPostRepository) GetThreadHistory(ctx context.Context, no int, board string) ([]*model.Post, error) {
	isop, err := mpr.IsOp(ctx, no, board)
	if err != nil {
		return nil, err
	}

	if !isop {
		return nil, model.NewBadRequest(fmt.Sprintf("post %v on board %v is not OP", no, board))
	}

	mpr.mutex.Lock()
	defer mpr.mutex.Unlock()
	var posts []*model.Post
	for _, p := range mpr.Posts[board] {
		if p.No == no || p.Parent == no {
			posts = append(posts, p)
		}
	}
	return posts, nil
}

func (mpr *memoryPostRepository) DeletePost(ctx context.Context, no int, board string) (bool, error) {
	mpr.mutex.Lock()
	defer mpr.mutex.Unlock()
	for _, p := range mpr.Posts[board] {
		if p.No == no {
			p.Com = "content deleted"
			return true, nil
		}
	}
	return false, model.NewNotFound("post", fmt.Sprintf("%d", no))
}

func (mpr *memoryPostRepository) IsOp(ctx context.Context, no int, board string) (bool, error) {
	mpr.mutex.Lock()
	defer mpr.mutex.Unlock()
	for _, p := range mpr.Posts[board] {
		if p.No == no {
			return p.Parent == 0, nil
		}
	}
	return false, model.NewNotFound("post", fmt.Sprintf("%d", no))
}
