package model

type BoardRepository interface {
	NewBoard(slug, name, descr string) (*Board, error)
	GetBoard(slug string) (*Board, error)
	GetBoardList() ([]*Board, error)
	IsBoardExists(slug string) (bool, error)
}

type PostRepository interface {
	NewPost(resto int, board, com string) (*Post, error)
	GetPost(no int, board string) (*Post, error)
	GetThreadHistory(no int, board string) ([]*Post, error)
	GetThreadList(board string) ([]*Post, error)
	DeletePost(no int, board string) (bool, error)
	IsOp(no int, board string) (bool, error)
}
