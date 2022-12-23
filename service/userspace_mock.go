package service

import (
	"context"
	"fmt"
	"time"

	"github.com/1ort/goimbo/model"
)

type UserspaceServiceMock struct {
}

// type UserspaceServiceConfig struct {
// 	PostRepository  model.PostRepository
// 	BoardRepository model.BoardRepository
// }

// func UserspaceServiceMock(c *UserspaceServiceConfig) model.Userspace {
// 	return &UserspaceService{
// 		PostRepository:  c.PostRepository,
// 		BoardRepository: c.BoardRepository,
// 	}
// }

func NewMockUserspace() model.Userspace {
	return &UserspaceServiceMock{}
}

func (u *UserspaceServiceMock) GetBoards(ctx context.Context) ([]*model.Board, error) {
	return []*model.Board{
		{Slug: "b", Name: "Board B", Descr: "Board B"},
		{Slug: "a", Name: "Board A", Descr: "Board A"},
		{Slug: "c", Name: "Board C", Descr: "Board C"},
	}, nil

}

func (u *UserspaceServiceMock) GetThread(ctx context.Context, board string, no int) (*model.Thread, error) {
	posts := make([]*model.Post, 10)
	for i := 0; i < 10; i++ {
		posts[i] = &model.Post{
			No:     i,
			Parent: 0,
			Board:  board,
			Com:    fmt.Sprintf("Post %d", i),
			Time:   time.Now().Unix(),
		}
	}
	return &model.Thread{
		OP:      posts[0],
		Replies: posts[1:],
	}, nil
}
func (u *UserspaceServiceMock) GetThreadPreview(ctx context.Context, board string, no int) (*model.ThreadPreview, error) {
	posts := make([]*model.Post, 10)
	for i := 0; i < 10; i++ {
		posts[i] = &model.Post{
			No:     i,
			Parent: 0,
			Board:  board,
			Com:    fmt.Sprintf("Post %d", i),
			Time:   time.Now().Unix(),
		}
	}
	return &model.ThreadPreview{
		OP:             posts[0],
		Replies:        10,
		OmittedReplies: 7,
		LastReplies:    posts[3:],
		LastModified:   time.Now().Unix(),
	}, nil
}
func (u *UserspaceServiceMock) GetBoardPage(ctx context.Context, board string, page int) (*model.BoardPage, error) {
	threads := make([]*model.ThreadPreview, 10)
	for i := 0; i < 10; i++ {
		posts := make([]*model.Post, 10)
		for j := 0; j < 10; j++ {
			posts[j] = &model.Post{
				No:     j,
				Parent: 0,
				Board:  board,
				Com:    fmt.Sprintf("Post %d", j),
				Time:   time.Now().Unix(),
			}
		}
		threads[i] = &model.ThreadPreview{
			OP:             posts[0],
			Replies:        10,
			OmittedReplies: 7,
			LastReplies:    posts[7:],
			LastModified:   time.Now().Unix(),
		}
	}
	return &model.BoardPage{
		Page:    page,
		Threads: threads,
	}, nil
}

func (u *UserspaceServiceMock) NewThread(ctx context.Context, board, com string) error {
	return nil
}
func (u *UserspaceServiceMock) Reply(ctx context.Context, board, com string, parent int) error {
	return nil
}
