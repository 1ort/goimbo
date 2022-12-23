package model

type Board struct {
	Slug  string `db:"slug" json:"slug"`
	Name  string `db:"name" json:"name"`
	Descr string `db:"descr" json:"descr"`
}

type Post struct {
	No     int    `db:"no" json:"no"`
	Parent int    `db:"parent" json:"parent"`
	Board  string `db:"board" json:"board"`
	Com    string `db:"com" json:"com"`
	Time   int64  `db:"time" json:"time"`
}

type Thread struct {
	OP      *Post   `json:"op"`
	Replies []*Post `json:"replies"`
}

type ThreadPreview struct {
	OP             *Post   `json:"op"`
	Replies        int     `json:"replies"`
	OmittedReplies int     `json:"ommited_posts"`
	LastReplies    []*Post `json:"last_replies"`
	LastModified   int64   `json:"last_modified"`
}

type BoardPage struct {
	Page    int              `json:"page"`
	Threads []*ThreadPreview `json:"threads"`
}
