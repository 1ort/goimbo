package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewBoard(pool *pgxpool.Pool, slug string, name string, descr string) error {
	new_board_query := `
	INSERT INTO board (slug, name, descr) VALUES
	($1, $2, $3)
	`
	_, err := pool.Query(context.Background(), new_board_query, slug, name, descr)
	if err != nil {
		return err
	}
	root_post_query := `
	INSERT INTO post (no, resto, board, com, time) VALUES
	(0, 0, $1, $2, NOW())
	`
	_, err = pool.Query(context.Background(), root_post_query, slug, "root_post")
	if err != nil {
		return err
	}
	return nil
}
