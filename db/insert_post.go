package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InsertPost(pool *pgxpool.Pool, resto int, board string, com string) error {
	query_template := `
	INSERT INTO post (no, resto, board, com, time) SELECT
	MAX (no) + 1, $1, $2, $3, NOW() FROM post
	WHERE (board=$2)
	`
	_, err := pool.Query(context.Background(), query_template, resto, board, com)
	return err
}

func IsOp(pool *pgxpool.Pool, board string, no int) bool {
	if no == 0 {
		return true
	}

	query_template := `
	SELECT EXISTS(SELECT 1 FROM post WHERE board=$1 and no=$2 and resto=0)
	`
	rows, query_err := pool.Query(context.Background(), query_template, board, no)
	if query_err != nil {
		fmt.Printf("%e", query_err)
		return false
	}
	defer rows.Close()
	var outs []bool
	for rows.Next() {
		var val bool
		read_err := rows.Scan(&val)
		if read_err != nil {
			fmt.Printf("%e", read_err)
			return false
		}
		outs = append(outs, val)
	}
	return outs[0]
}
