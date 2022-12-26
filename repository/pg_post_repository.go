package repository

import (
	"context"

	"github.com/1ort/goimbo/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var schema = `CREATE TABLE IF NOT EXISTS posts (
	no        INT         NOT NULL,
	board     TEXT        NOT NULL,
	parent    INT         NOT NULL,
	com       TEXT        NOT NULL,
	time      TIMESTAMPTZ NOT NULL,
	UNIQUE (no, board),
	PRIMARY KEY (no, board)
  );
`

type PgPostRepository struct {
	connPool *pgxpool.Pool
}

type PgPostRepoConfig struct {
	Pool *pgxpool.Pool
}

func NewPGPostRepository(cfg *PgPostRepoConfig) model.PostRepository {
	p := &PgPostRepository{
		connPool: cfg.Pool,
	}
	_, err := p.connPool.Exec(context.Background(), schema)
	if err != nil {
		panic(err)
	}
	return p
}

func (p *PgPostRepository) NewPost(ctx context.Context, board, com string, parent int) (*model.Post, error) {
	query_template :=
		`INSERT INTO posts (no, board, parent, com, time)
		VALUES (
		  (SELECT COALESCE(MAX(no) + 1, 1) FROM posts WHERE board = $1),
		  $1,
		  $2,
		  $3,
		  NOW()
		)
		RETURNING *
		`
	var post = model.Post{}
	row := p.connPool.QueryRow(ctx, query_template, board, parent, com)
	err := row.Scan(&post.No, &post.Board, &post.Parent, &post.Com, &post.Time)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (p *PgPostRepository) GetSingle(ctx context.Context, board string, no int) (*model.Post, error) {
	query_template :=
		`SELECT * FROM posts
		WHERE board = $1 AND no = $2
		LIMIT 1`

	var post = model.Post{}
	row := p.connPool.QueryRow(ctx, query_template, board, no)
	err := row.Scan(&post.No, &post.Board, &post.Parent, &post.Com, &post.Time)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

/*reverse = от новых к старым*/
func (p *PgPostRepository) GetMultiple(ctx context.Context, board string, parent int, skip, limit int, reverse, sort_by_last_modified bool) ([]*model.Post, error) {
	if limit == 0 {
		limit = 100_000
	}
	base_query_template :=
		`SELECT * FROM posts
		WHERE board = $1 AND parent = $2
		order by 
			case when $5 then time end desc
  		, case when not $5 then time end asc
		LIMIT $3 OFFSET $4
		`
	last_modified_query_template :=
		`SELECT p1.no, p1.board, p1.parent, p1.com, p1.time
		FROM posts p1
		LEFT JOIN posts p2 ON p2.parent = p1.no AND p2.board = p1.board
		WHERE p1.board = $1 AND p1.parent = $2
		GROUP BY p1.no, p1.board, p1.parent, p1.com, p1.time
		order by 
			case when $5 then COALESCE(MAX(p2.no), p1.no) end desc
  		, case when not $5 then COALESCE(MAX(p2.no), p1.no) end asc
		LIMIT $3 OFFSET $4
		`
	var query_template string
	if sort_by_last_modified {
		query_template = last_modified_query_template
	} else {
		query_template = base_query_template
	}

	rows, err := p.connPool.Query(ctx, query_template, board, parent, limit, skip, reverse)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*model.Post
	posts, err = pgx.CollectRows(rows,
		func(row pgx.CollectableRow) (*model.Post, error) {
			var post = &model.Post{}
			err := row.Scan(&post.No, &post.Board, &post.Parent, &post.Com, &post.Time)
			return post, err
		})
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (p *PgPostRepository) Count(ctx context.Context, board string, parent int) (int, error) {
	query_template :=
		`SELECT COUNT(*) FROM posts
	WHERE board = $1 AND parent = $2
	`
	var c int
	row := p.connPool.QueryRow(ctx, query_template, board, parent)
	err := row.Scan(&c)
	if err != nil {
		return 0, err
	}
	return c, nil
}
