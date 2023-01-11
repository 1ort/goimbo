package model

import (
	"context"
	"io"
	"mime/multipart"

	"github.com/google/uuid"
)

// repository layer
type BoardRepository interface {
	NewBoard(ctx context.Context, slug, name, descr string) (*Board, error)
	GetBoard(ctx context.Context, slug string) (*Board, error)
	GetBoardList(ctx context.Context) ([]*Board, error)
}

type PostRepository interface {
	NewPost(ctx context.Context, board, com string, parent int) (*Post, error)
	GetSingle(ctx context.Context, board string, no int) (*Post, error)
	GetMultiple(ctx context.Context, board string, parent int, skip, limit int, reverse, sortByLastModified bool) ([]*Post, error)
	Count(ctx context.Context, board string, parent int) (int, error)
}

type AttachmentRepository interface {
	SaveAttachments(ctx context.Context, board string, post int, attachments []*Attachment) ([]*Attachment, error)
	GetMultiple(ctx context.Context, board string, post int) ([]*Attachment, error)
	GetSingle(ctx context.Context, UUID uuid.UUID) (*Attachment, error)
}

// service layer
type Userspace interface {
	GetBoards(ctx context.Context) ([]*Board, error)
	GetBoard(ctx context.Context, slug string) (*Board, error)
	GetThread(ctx context.Context, board string, no int) (*Thread, error)
	//GetThreadPreview(ctx context.Context, board string, no int) (*ThreadPreview, error)
	GetBoardPage(ctx context.Context, board string, page int) (*BoardPage, error)
	NewThread(ctx context.Context, board, com string) (*Post, error)
	Reply(ctx context.Context, board, com string, parent int) (*Post, error)
}

type AttachmentService interface {
	AttachFromFileHeaders(ctx context.Context, files []*multipart.FileHeader, post *Post) ([]*Attachment, error)
	// GetFromPost(ctx context.Context, post *Post) ([]*Attachment, Error)
	// WriteContent(uuid uuid.UUID)
	// WritePreview(uuid uuid.UUID)
}

type Adminspace interface {
}

type Captcha interface {
	New() (string, error)
	Verify(id, solution string) (bool, error)
	//Write captcha task
	Write(w io.Writer, id string) error
}

type MultipartFileSaver interface {
}
