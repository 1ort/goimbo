package service

import (
	"context"
	"fmt"
	"math"
	"strconv"

	"github.com/1ort/goimbo/model"
)

type UserspaceService struct {
	PostRepository  model.PostRepository
	BoardRepository model.BoardRepository
	ThreadsPerPage  int
	PreviewPosts    int
}

type UserspaceServiceConfig struct {
	PostRepository  model.PostRepository
	BoardRepository model.BoardRepository
}

func NewUserspaceService(c *UserspaceServiceConfig) model.Userspace {
	return &UserspaceService{
		PostRepository:  c.PostRepository,
		BoardRepository: c.BoardRepository,
		ThreadsPerPage:  3,
		PreviewPosts:    3,
	}
}

func (u *UserspaceService) GetBoard(ctx context.Context, slug string) (*model.Board, error) {
	return u.BoardRepository.GetBoard(ctx, slug)
}

func (u *UserspaceService) GetBoards(ctx context.Context) ([]*model.Board, error) {
	return u.BoardRepository.GetBoardList(ctx)
}

func (u *UserspaceService) GetThread(ctx context.Context, board string, no int) (*model.Thread, error) {
	if _, err := u.BoardRepository.GetBoard(ctx, board); err != nil {
		return nil, model.NewNotFound("board", board)
	}
	op, err := u.getOP(ctx, board, no)
	if err != nil {
		return nil, err
	}
	replies, err := u.PostRepository.GetMultiple(ctx, board, no, 0, 0, false, false)
	if err != nil {
		return nil, err
	}
	return &model.Thread{
		OP:      op,
		Replies: replies,
	}, nil
}

func (u *UserspaceService) GetThreadPreview(ctx context.Context, board string, no int) (*model.ThreadPreview, error) {
	if _, err := u.BoardRepository.GetBoard(ctx, board); err != nil {
		return nil, model.NewNotFound("board", board)
	}
	op, err := u.getOP(ctx, board, no)
	if err != nil {
		return nil, err
	}
	return u.getThreadPreviewByOp(ctx, op)
}

func (u *UserspaceService) getThreadPreviewByOp(ctx context.Context, op *model.Post) (*model.ThreadPreview, error) {
	repliesCount, err := u.PostRepository.Count(ctx, op.Board, op.No)
	if err != nil {
		return nil, err
	}
	ommitedPosts := repliesCount - u.PreviewPosts
	if ommitedPosts < 0 {
		ommitedPosts = 0
	}
	replies, err := u.PostRepository.GetMultiple(ctx, op.Board, op.No, ommitedPosts, 0, false, false)
	if err != nil {
		return nil, err
	}
	LastModified := op.Time
	if len(replies) > 0 {
		LastModified = replies[len(replies)-1].Time
	}
	return &model.ThreadPreview{
		OP:             op,
		TotalReplies:   repliesCount,
		OmittedReplies: ommitedPosts,
		LastReplies:    replies,
		LastModified:   LastModified,
	}, nil
}

func (u *UserspaceService) GetBoardPage(ctx context.Context, board string, page int) (*model.BoardPage, error) {
	if _, err := u.BoardRepository.GetBoard(ctx, board); err != nil {
		return nil, model.NewNotFound("board", board)
	}
	var limit, offset, totalThreads, totalPages int
	var threads []*model.ThreadPreview
	limit = u.ThreadsPerPage
	offset = u.ThreadsPerPage * page
	totalThreads, err := u.PostRepository.Count(ctx, board, 0)
	if err != nil {
		return nil, err
	}
	totalPages = int(math.Ceil(float64(totalThreads) / float64(u.ThreadsPerPage)))
	if page >= totalPages && totalPages != 0 {
		return nil, model.NewNotFound("page", strconv.Itoa(page))
	}
	ops, err := u.PostRepository.GetMultiple(ctx, board, 0, offset, limit, true, true)
	if err != nil {
		return nil, err
	}
	var thread *model.ThreadPreview
	for _, op := range ops {
		thread, err = u.getThreadPreviewByOp(ctx, op)
		if err != nil {
			return nil, err
		}
		threads = append(threads, thread)
	}
	return &model.BoardPage{
		Page: &model.PageValue{
			CurrentPage: page,
			TotalPages:  totalPages,
		},
		Threads: threads,
	}, nil
}

func (u *UserspaceService) NewThread(ctx context.Context, board, com string) (*model.Post, error) {
	return u.PostRepository.NewPost(ctx, board, com, 0)
}

func (u *UserspaceService) Reply(ctx context.Context, board, com string, parent int) (*model.Post, error) {
	_, err := u.getOP(ctx, board, parent)
	if err != nil {
		return nil, err
	}
	return u.PostRepository.NewPost(ctx, board, com, parent)
}

func (u *UserspaceService) getOP(ctx context.Context, board string, no int) (*model.Post, error) {
	op, err := u.PostRepository.GetSingle(ctx, board, no)
	if err != nil {
		return nil, err
	}
	if op.Parent != 0 {
		return nil, model.NewBadRequest(fmt.Sprintf("post #%v is not OP", no))
	}
	return op, nil
}
