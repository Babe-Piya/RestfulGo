package repository

import (
	// "log"
	"restfulGo/entity"
)

type TodoRepo struct {
	TodoList []*entity.Todo
	InitId int64 
}

func CreateRepo() TodoRepoInterface {
	return &TodoRepo{InitId:1}
}

func (tr *TodoRepo) Add(todo *entity.Todo) (int64, error) {
	todo.Id = tr.InitId
	tr.TodoList = append(tr.TodoList, todo)
	tr.InitId++
	// log.Print("todo list post here", tr.TodoList)

	return todo.Id, nil
}

func (tr *TodoRepo) Get() []entity.Todo {
	var todoList []entity.Todo
	for _, todo := range tr.TodoList {
		// log.Println("todolist get is ", todo)
		todoList = append(todoList, *todo)
	}
	return todoList

}

func (tr *TodoRepo) Del(Id int64) (entity.Todo, error) {
	var todoList *entity.Todo
	for index:= range tr.TodoList {
		if tr.TodoList[index].Id == Id {
			todoList = tr.TodoList[index]
			if tr.TodoList[len(tr.TodoList)-1].Id == Id {
				tr.TodoList[index] = todoList
			} else {
				copy(tr.TodoList[index:],tr.TodoList[index+1:])
				tr.TodoList[len(tr.TodoList)-1] = todoList
			}
			tr.TodoList = tr.TodoList[:len(tr.TodoList)-1]
			break
		}
	}
	return *todoList, nil
}

func (tr *TodoRepo) Update(Id int64, todo *entity.Todo) (entity.Todo,error){
	var todoList *entity.Todo
	for index:= range tr.TodoList {
		if tr.TodoList[index].Id == Id {
			tr.TodoList[index] = todo
			todoList  = tr.TodoList[index]
			break
		}
	}
	return *todoList, nil
}
