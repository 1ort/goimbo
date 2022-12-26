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
	op, err := u.getOP(ctx, board, no)
	if err != nil {
		return nil, err
	}
	return u.getThreadPreviewByOp(ctx, op)
}

func (u *UserspaceService) getThreadPreviewByOp(ctx context.Context, op *model.Post) (*model.ThreadPreview, error) {
	replies_count, err := u.PostRepository.Count(ctx, op.Board, op.No)
	if err != nil {
		return nil, err
	}
	ommited_posts := replies_count - u.PreviewPosts
	if ommited_posts < 0 {
		ommited_posts = 0
	}
	replies, err := u.PostRepository.GetMultiple(ctx, op.Board, op.No, ommited_posts, 0, false, false)
	if err != nil {
		return nil, err
	}
	last_modified := op.Time
	if len(replies) > 0 {
		last_modified = replies[len(replies)-1].Time
	}
	return &model.ThreadPreview{
		OP:             op,
		TotalReplies:   replies_count,
		OmittedReplies: ommited_posts,
		LastReplies:    replies,
		LastModified:   last_modified,
	}, nil
}

/*TODO: Зафакапил, нужно было сначала сортировать, а потом делить на страницы*/
func (u *UserspaceService) GetBoardPage(ctx context.Context, board string, page int) (*model.BoardPage, error) {
	var limit, offset, total_threads, total_pages int
	var threads []*model.ThreadPreview
	limit = u.ThreadsPerPage
	offset = u.ThreadsPerPage * page
	total_threads, err := u.PostRepository.Count(ctx, board, 0)
	if err != nil {
		return nil, err
	}
	total_pages = int(math.Ceil(float64(total_threads) / float64(u.ThreadsPerPage)))
	if page >= total_pages && total_pages != 0 {
		//fmt.Printf("page: %v, total_pages")
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
	// //sort threads by last modified value
	// sort.Slice(threads, func(i, j int) bool {
	// 	return threads[i].LastModified.After(threads[j].LastModified)
	// })
	return &model.BoardPage{
		Page: &model.PageValue{
			CurrentPage: page,
			TotalPages:  total_pages,
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
