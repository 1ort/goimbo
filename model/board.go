package model

type Board struct {
	Slug  string `db:"slug" json:"slug"`
	Name  string `db:"name" json:"name"`
	Descr string `db:"descr" json:"descr"`
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
