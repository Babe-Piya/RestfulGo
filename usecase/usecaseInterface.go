package usecase

import (
	"restfulGo/entity"
)

type TodoUsecaseInterface interface{
	Add(todo *entity.Todo) (int64,error)
	Get() []entity.Todo
	// Del(Id int64) (entity.Todo,error)
	// Update(Id int64, todo *entity.Todo) (entity.Todo,error)
}