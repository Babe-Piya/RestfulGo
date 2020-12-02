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

			return  uc.repo.Del(Id)

}

func (uc TodoUsecase) Update(Id int64,todo map[string]interface{}) (entity.Todo,error){
	var todoList entity.Todo
			 uc.repo.Update(Id,todo)

			return todoList,nil

}
