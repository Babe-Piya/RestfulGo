package handler

import (
	"restfulGo/usecase"
)

type Handler struct {
	uc usecase.TodoUsecaseInterface
}

type Req struct {
	Content string `json:"content"`
	Title   string `json:"title" binding:"required"`
}

type ReqPut struct {
	Content string `json:"content"`
	Title   string `json:"title"`
}

func CreateHandler(ucm usecase.TodoUsecaseInterface) Handler {
	return Handler{uc: ucm}
}

