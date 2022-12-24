package repository

import (
	"context"

	"github.com/1ort/goimbo/model"
)

type PgPostRepository struct {
}

type PgPostRepoConfig struct {
	//connection
}

func (p *PgPostRepository) NewPost(ctx context.Context, board, com string, parent int) error {
	return nil
}

func (p *PgPostRepository) GetSingle(ctx context.Context, board string, no int) (*model.Post, error) {
	return nil, nil
}

func (p *PgPostRepository) GetMultiple(ctx context.Context, board string, parent int, skip, limit int) ([]*model.Post, error) {
	return nil, nil
}
