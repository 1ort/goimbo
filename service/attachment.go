package service

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/1ort/goimbo/model"
	"github.com/google/uuid"
)

type attachmentService struct {
	repo              model.AttachmentRepository
	folder            string
	allowedExtensions []string
	maxSize           float32
	maxAttachments    int
	minAttachments    int
}

type AttachmentServiceConfig struct {
	Repo              model.AttachmentRepository
	Folder            string
	AllowedExtensions []string
	MaxSize           float32
	MaxAttachments    int
	MinAttachments    int
}

func NewAttachmentService(cfg *AttachmentServiceConfig) model.AttachmentService {
	return &attachmentService{
		repo:              cfg.Repo,
		folder:            cfg.Folder,
		allowedExtensions: cfg.AllowedExtensions,
		maxSize:           cfg.MaxSize,
		maxAttachments:    cfg.MaxAttachments,
		minAttachments:    cfg.MinAttachments,
	}
}

func (as *attachmentService) allowedMaxSize(files []*multipart.FileHeader) bool {
	var size float32
	for _, file := range files {
		size += float32(file.Size)
	}
	return size <= as.maxSize
}

func (as *attachmentService) saveFile(file *multipart.FileHeader, name string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	dst := filepath.Join(as.folder, name)
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

func (as *attachmentService) AttachFromFileHeaders(ctx context.Context, files []*multipart.FileHeader, post *model.Post) ([]*model.Attachment, error) {
	if len(files) > as.maxAttachments || len(files) < as.minAttachments {
		return nil, model.NewBadRequest("Invalid number of files")
	}
	if !as.allowedMaxSize(files) {
		return nil, model.NewBadRequest("Attachments maximum size exceeded")
	}
	var attachments []*model.Attachment
	for _, file := range files {
		attachment := &model.Attachment{
			UUID:     uuid.New(),
			Filename: file.Filename,
			FileSize: int(file.Size),
		}
		extension := filepath.Ext(attachment.Filename)
		name := attachment.UUID.String() + extension
		err := as.saveFile(file, name)
		if err != nil {
			return nil, model.NewBadRequest("File save error")
		}
		attachments = append(attachments, attachment)
	}
	saved, err := as.repo.SaveAttachments(ctx, post.Board, post.No, attachments)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	return saved, nil
}
