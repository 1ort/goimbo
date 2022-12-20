package model

import (
	"context"
)

type BoardRepository interface {
	NewBoard(ctx context.Context, slug, name, descr string) (*Board, error)
	GetBoard(ctx context.Context, slug string) (*Board, error)
	GetBoardList(ctx context.Context) ([]*Board, error)
	IsBoardExists(ctx context.Context, slug string) (bool, error)
}

type PostRepository interface {
	NewPost(ctx context.Context, resto int, board, com string) (*Post, error)
	GetPost(ctx context.Context, no int, board string) (*Post, error)
	GetThreadHistory(ctx context.Context, no int, board string) ([]*Post, error)
	GetThreadList(ctx context.Context, board string) ([]*Post, error)
	DeletePost(ctx context.Context, no int, board string) (bool, error)
	IsOp(ctx context.Context, no int, board string) (bool, error)
}
