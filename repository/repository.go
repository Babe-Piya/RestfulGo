package repository

import (
	"log"
	"restfulGo/entity"
)

type TodoRepo struct{
	TodoList []*entity.Todo
}

func CreateRepo() TodoRepoInterface{
	return &TodoRepo{}
}

func (tr *TodoRepo) Add(todo *entity.Todo) (int64,error){
	if len(tr.TodoList) == 0 {
		todo.Id  = 1
		} else {
			todo.Id = tr.TodoList[len(tr.TodoList)-1].Id+1

		}
		tr.TodoList = append(tr.TodoList,todo)
		log.Print("todo list here",tr.TodoList)
	
	return todo.Id,nil
}

func (tr *TodoRepo) Get() []entity.Todo {
	var todoList []entity.Todo
	for _,todo := range tr.TodoList {
		log.Println("todo is ", todo)
		todoList = append(todoList,*todo)
	}
	return todoList
}