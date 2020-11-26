package usecase

import (
	"restfulGo/entity"
	"restfulGo/repository"
	"errors"
)

type TodoUsecase struct {
	repo repository.TodoRepoInterface
}

func CreateUsecase(repo repository.TodoRepoInterface) TodoUsecaseInterface {
	return TodoUsecase{repo: repo}
}

func (uc TodoUsecase) Add(todo *entity.Todo) (int64, error) {
	id, err := uc.repo.Add(todo)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (uc TodoUsecase) Get() []entity.Todo {
	return uc.repo.Get()
}

func (uc TodoUsecase) Del(Id int64) (entity.Todo, error) {
	var todo entity.Todo
	findID := uc.repo.Get()
	for index := range findID {
		if findID[index].Id == Id {
			return  uc.repo.Del(Id)
		} 
	} 
	return todo,errors.New("don't have Id")
}

func (uc TodoUsecase) Update(Id int64,todo *entity.Todo) (entity.Todo,error){
	findTodoList := uc.repo.Get()
	for index,value := range findTodoList {
		if findTodoList[index].Id == Id {
			todo.Id = Id
			if todo.Content == "" {
				todo.Content = value.Content
			}
			if todo.Title == "" {
				todo.Title = value.Title
			}
			return uc.repo.Update(Id,todo)
		} 
	} 
	return *todo,errors.New("don't have Id")
}
