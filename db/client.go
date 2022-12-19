package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitPool(database_url string) *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), database_url)
	if err != nil {
		panic(fmt.Sprintf("Unable to create connection pool: %v\n", err))
	}
	return dbpool
}

func InitDatabase(pool *pgxpool.Pool) {
	for i := 0; i < len(Schema); i++ {
		_, err := pool.Query(context.Background(), Schema[i])
		if err != nil {
			panic(fmt.Sprintf("Unable to init database: %v\n", err))
		}
	}
}
