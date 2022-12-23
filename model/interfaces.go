package model

import (
	"context"
)

/*
TODO:
Refactor PostRepository to meet the needs of the Userspace service

Remove IsBoardExists in BoardRepository
IsBoardExists can be implemented inside a service via GetBoard
*/

type BoardRepository interface {
	NewBoard(ctx context.Context, slug, name, descr string) (*Board, error)
	GetBoard(ctx context.Context, slug string) (*Board, error)
	GetBoardList(ctx context.Context) ([]*Board, error)
	IsBoardExists(ctx context.Context, slug string) (bool, error)
}

type PostRepository interface {
	NewPost(ctx context.Context, parent int, board, com string) (*Post, error)
	GetPost(ctx context.Context, no int, board string) (*Post, error)
	GetThreadHistory(ctx context.Context, no int, board string) ([]*Post, error)
	GetThreadList(ctx context.Context, board string) ([]*Post, error)
	DeletePost(ctx context.Context, no int, board string) (bool, error)
	IsOp(ctx context.Context, no int, board string) (bool, error)
}

type PostRepositoryRef interface {
	NewPost(ctx context.Context, board, com string, parent int) error
	//DeletePost(ctx context.Context, board, com string)
	GetSingle(ctx context.Context, board string, no int) (*Post, error)
	GetMultiple(ctx context.Context, board string, parent int, skip, limit int) ([]*Post, error)
}

type Userspace interface {
	GetBoards(ctx context.Context) ([]*Board, error)
	GetThread(ctx context.Context, board string, no int) (*Thread, error)
	//GetThreadPreview(ctx context.Context, board string, no int) (*ThreadPreview, error)
	GetBoardPage(ctx context.Context, board string, page int) (*BoardPage, error)
	NewThread(ctx context.Context, board, com string) error
	Reply(ctx context.Context, board, com string, parent int) error
}

type Adminspace interface {
}
