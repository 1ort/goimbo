package model

type Board struct {
	Slug  string `db:"slug" json:"slug"`
	Name  string `db:"name" json:"name"`
	Descr string `db:"descr" json:"descr"`
}

type BoardRepository interface {
	NewBoard(slug, name, descr string) (b *Board, err error)
	GetBoard(slug string) (b *Board, err error)
	GetBoardList() (l []*Board, err error)
	IsExists(slug string) (exists bool, err error)
}

/*
schema:
CREATE TABLE IF NOT EXISTS board (
	slug text UNIQUE PRIMARY KEY,
	name text,
	descr text
)

new_board_query:
INSERT INTO board (slug, name, descr) VALUES
($1, $2, $3)

IsExists_query:
SELECT EXISTS(SELECT 1 FROM board WHERE slug=$1)
*/
