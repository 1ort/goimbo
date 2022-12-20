package model

type Board struct {
	Slug  string `db:"slug" json:"slug"`
	Name  string `db:"name" json:"name"`
	Descr string `db:"descr" json:"descr"`
}

type Post struct {
	No    int    `db:"no" json:"no"`
	Resto int    `db:"resto" json:"resto"`
	Board string `db:"board" json:"board"`
	Com   string `db:"com" json:"com"`
	Time  int    `db:"time" json:"time"`
}

type ThreadInfo struct {
	OP           int `json:"no"`
	LastModified int `json:"last_modified"`
	Replies      int `json:"replies"`
}

type ThreadListPage struct {
	Page    int          `json:"page"`
	Threads []ThreadInfo `json:"threads"`
}

type CatalogThread struct {
	OmittedPosts int `json:"omitted_posts"`
	Replies      int `json:"replies"`
	Post
	LastReplies  []Post `json:"last_replies"`
	LastModified int    `json:"last_modified"`
}

type CatalogPage struct {
	Page    int             `json:"page"`
	Threads []CatalogThread `json:"threads"`
}

type ThreadPage struct {
	Page    int      `json:"page"`
	Threads [][]Post `json:"threads"`
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
