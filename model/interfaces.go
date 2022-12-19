package model

type BoardRepository interface {
	NewBoard(slug, name, descr string) (b *Board, err error)
	GetBoard(slug string) (b *Board, err error)
	GetBoardList() (l []*Board, err error)
	IsExists(slug string) (exists bool, err error)
}

type PostRepository interface {
	NewPost(resto int, board, com string) (p *Post, err error)
	GetPost(no int, board string) (p *Post, err error)
	GetThreadHistory(no int, board string)
	DeletePost(no int, board string) (deleted bool, err error)
	IsOp(no int, board string) (isop bool, err error)
}
