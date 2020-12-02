package repository

import (
	"restfulGo/entity"
)

type TodoRepoInterface interface {
	Add(todo *entity.Todo) (int64, error)
	Get() []entity.Todo
	Del(Id int64) (entity.Todo, error)
	Update(Id int64, todo map[string]interface{}) (entity.Todo,error)
}
