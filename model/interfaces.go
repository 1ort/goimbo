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

type Userspace interface {
	Boards(ctx context.Context) ([]*Board, error)
	Threads(ctx context.Context, board string) ([]*ThreadListPage, error)
	Catalog(ctx context.Context, board string) ([]*CatalogPage, error)
	Index(ctx context.Context, board string, page int) (*ThreadPage, error)
	Thread(ctx context.Context, board string, op int) ([]*Post, error)
	NewPost(ctx context.Context, board string, parent int, com string) (*Post, error)
}

type Adminspace interface {
}
