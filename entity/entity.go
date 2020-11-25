package entity

import "time"

type Todo struct {
	Id int64
	Title string
	Content string
	IsDone bool
	CreateAt time.Time
}