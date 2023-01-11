package repository

import (
	"context"

	"github.com/1ort/goimbo/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var attachemtSchema = `CREATE TABLE IF NOT EXISTS attachments (
	uuid      TEXT      UNIQUE PRIMARY KEY,
	filename  TEXT        NOT NULL,
	filesize  INT         NOT NULL,
	board     TEXT        NOT NULL,
	post      INT         NOT NULL
  );`

type pgAttachmentRepository struct {
	connPool *pgxpool.Pool
}

type PGAttachmentRepoConfig struct {
	ConnPool *pgxpool.Pool
}

func NewPGAttachmentRepository(cfg *PGAttachmentRepoConfig) model.AttachmentRepository {
	p := &pgAttachmentRepository{
		connPool: cfg.ConnPool,
	}
	_, err := p.connPool.Exec(context.Background(), attachemtSchema)
	if err != nil {
		panic(err)
	}
	return p
}

func (r *pgAttachmentRepository) SaveAttachments(ctx context.Context, board string, post int, attachments []*model.Attachment) ([]*model.Attachment, error) {
	queryTemplate :=
		`INSERT INTO attachments (uuid, filename, filesize, board, post)
		VALUES (
		  $1,
		  $2,
		  $3,
		  $4,
		  $5
		)
		`
	for _, attachment := range attachments {
		_, err := r.connPool.Exec(ctx, queryTemplate, attachment.UUID, attachment.Filename, attachment.FileSize, board, post)
		if err != nil {
			return nil, err
		}
	}
	return attachments, nil
}

func (r *pgAttachmentRepository) GetMultiple(ctx context.Context, board string, post int) ([]*model.Attachment, error) {
	queryTemplate :=
		`SELECT * FROM attachments
		WHERE board = $1 AND post = $2
		`
	rows, err := r.connPool.Query(ctx, queryTemplate, board, post)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var attachments []*model.Attachment
	attachments, err = pgx.CollectRows(rows,
		func(row pgx.CollectableRow) (*model.Attachment, error) {
			var attachment = &model.Attachment{}
			err := row.Scan(&attachment.UUID, &attachment.Filename, &attachment.FileSize)
			return attachment, err
		})
	if err != nil {
		return nil, err
	}
	return attachments, nil
}

func (r *pgAttachmentRepository) GetSingle(ctx context.Context, UUID uuid.UUID) (*model.Attachment, error) {
	queryTemplate :=
		`SELECT * FROM attachments
		WHERE uuid = $1
		LIMIT 1`

	var attachment = model.Attachment{}
	row := r.connPool.QueryRow(ctx, queryTemplate, UUID)
	err := row.Scan(&attachment.UUID, &attachment.Filename, &attachment.FileSize)
	if err != nil {
		return nil, err
	}
	return &attachment, nil
}
