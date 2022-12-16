package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPost(pool *pgxpool.Pool, resto string, board string, com string) error {
	query_template := `
	INSERT INTO post (no, resto, board, com, time) SELECT
	(MAX (no) + 1, $1, $2, $3, NOW()) FROM post
	WHERE (board=$2)
	`
	_, err := pool.Query(context.Background(), query_template, resto, board, com)
	return err
}
