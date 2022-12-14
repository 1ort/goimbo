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
}

type PostRepository interface {
	NewPost(ctx context.Context, board, com string, parent int) (*Post, error)
	GetSingle(ctx context.Context, board string, no int) (*Post, error)
	GetMultiple(ctx context.Context, board string, parent int, skip, limit int, reverse, sortByLastModified bool) ([]*Post, error)
	Count(ctx context.Context, board string, parent int) (int, error)
}

type Userspace interface {
	GetBoards(ctx context.Context) ([]*Board, error)
	GetBoard(ctx context.Context, slug string) (*Board, error)
	GetThread(ctx context.Context, board string, no int) (*Thread, error)
	//GetThreadPreview(ctx context.Context, board string, no int) (*ThreadPreview, error)
	GetBoardPage(ctx context.Context, board string, page int) (*BoardPage, error)
	NewThread(ctx context.Context, board, com string) (*Post, error)
	Reply(ctx context.Context, board, com string, parent int) (*Post, error)
}

type Adminspace interface {
}
