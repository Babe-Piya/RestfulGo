package usecase

import (
	"restfulGo/entity"
	"restfulGo/repository"
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
	todoList, err := uc.repo.Del(Id)
	return todoList, err
}

// func (uc TodoUsecase) Update(Id int64,todo *entity.Todo) (entity.Todo,error){

// 	return entity.Todo{},nil
// }
