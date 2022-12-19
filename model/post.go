package model

type Post struct {
	No    int32  `db:"no" json:"no"`
	Resto int32  `db:"resto" json:"resto"`
	Board string `db:"board" json:"board"`
	Com   string `db:"com" json:"com"`
	Time  int32  `db:"time" json:"time"`
}

type PostRepository interface {
	NewPost(resto int, board, com string) (p *Post, err error)
	GetPost(no int, board string) (p *Post, err error)
	GetThreadHistory(no int, board string)
	DeletePost(no int, board string) (deleted bool, err error)
	IsOp(no int, board string) (isop bool, err error)
}

/*
schema:
CREATE TABLE IF NOT EXISTS post (
	no INT NOT NULL,
	resto INT NOT NULL,
	board text NOT NULL,
	com text NOT NULL,
	time TIMESTAMP,
	UNIQUE (no, board),
	FOREIGN KEY (board) REFERENCES board(slug)
)

NewPostQuery:
INSERT INTO post (no, resto, board, com, time) SELECT
MAX (no) + 1, $1, $2, $3, NOW() FROM post
WHERE (board=$2)

IsOpQuery:
SELECT EXISTS(SELECT 1 FROM post WHERE board=$1 and no=$2 and resto=0)
*/
