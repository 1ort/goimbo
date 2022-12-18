package model

type Board struct {
	Slug  string `db:"slug" json:"slug"`
	Name  string `db:"name" json:"name"`
	Descr string `db:"descr" json:"descr"`
}
