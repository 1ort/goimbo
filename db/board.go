package db

import (
	"context"
	"fmt"

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

func BoardExists(pool *pgxpool.Pool, slug string) (bool, error) {
	query_template := `
	SELECT EXISTS(SELECT 1 FROM board WHERE slug=$1)
	`
	rows, query_err := pool.Query(context.Background(), query_template, slug)
	defer rows.Close()
	if query_err != nil {
		fmt.Printf("%e", query_err)
		return false, query_err
	}
	var outs []bool
	for rows.Next() {
		var val bool
		read_err := rows.Scan(&val)
		if read_err != nil {
			fmt.Printf("%e", read_err)
			return false, read_err
		}
		outs = append(outs, val)
	}
	return outs[0], nil
}
