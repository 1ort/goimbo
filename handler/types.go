package handler

// TODO: bindall() function
type PostRequest struct {
	Com string `form:"text" binding:"required"`
}

type BoardRequest struct {
	Board string `uri:"board" binding:"required"`
}

type BoardPageRequest struct {
	BoardRequest
	Page int `uri:"page" binding:"gte=0"`
}

type ThreadPageRequest struct {
	BoardRequest
	Thread int `uri:"thread" binding:"gte=0"`
}
