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
		{Slug: "b", Name: "Board B", Descr: "Board B description"},
		{Slug: "a", Name: "Board A", Descr: "Board A description"},
		{Slug: "c", Name: "Board C", Descr: "Board C description"},
	}, nil
}

func (u *UserspaceServiceMock) GetBoard(ctx context.Context, slug string) (*model.Board, error) {
	return &model.Board{
		Slug: "b", Name: "Board B", Descr: "Board B description",
	}, nil
}

func (u *UserspaceServiceMock) GetThread(ctx context.Context, board string, no int) (*model.Thread, error) {
	posts := make([]*model.Post, 15)
	for i := 0; i < 15; i++ {
		posts[i] = &model.Post{
			No:     i + 5,
			Parent: 5,
			Board:  board,
			Com:    fmt.Sprintf("Post %d Lorem ipsum dolor sit amet.", i),
			Time:   time.Now(),
		}
	}
	posts[0].Parent = 0
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
			Com:    fmt.Sprintf("Post %d Lorem ipsum dolor sit amet.", i),
			Time:   time.Now(),
		}
	}
	return &model.ThreadPreview{
		OP:             posts[0],
		TotalReplies:   10,
		OmittedReplies: 7,
		LastReplies:    posts[3:],
		LastModified:   time.Now(),
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
				Com:    fmt.Sprintf("Post %d Lorem ipsum dolor sit amet.", j),
				Time:   time.Now(),
			}
		}
		threads[i] = &model.ThreadPreview{
			OP:             posts[0],
			TotalReplies:   10,
			OmittedReplies: 7,
			LastReplies:    posts[7:],
			LastModified:   time.Now(),
		}
	}
	return &model.BoardPage{
		Page: &model.PageValue{
			CurrentPage: page,
			TotalPages:  15,
		},
		Threads: threads,
	}, nil
}

func (u *UserspaceServiceMock) NewThread(ctx context.Context, board, com string) error {
	return nil
}
func (u *UserspaceServiceMock) Reply(ctx context.Context, board, com string, parent int) error {
	return nil
}
