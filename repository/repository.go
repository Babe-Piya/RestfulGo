package repository

import (
	"strings"
	"fmt"
	"errors"
	"database/sql"
	"log"
	"restfulGo/entity"
	"github.com/gobeam/stringy"
)

type TodoRepo struct {
	DB *sql.DB
}

func CreateRepo(db *sql.DB) TodoRepoInterface {
	return &TodoRepo{DB:db}
}

func (tr *TodoRepo) Add(todo *entity.Todo) (int64, error) {
	stm,err := tr.DB.Prepare("INSERT INTO todo.todo (TITLE,CONTENT) VALUE (?,?)")
	
	if err!= nil {
		return 0,errors.New("error")
	}
	defer stm.Close()
	result,err := stm.Exec(&todo.Title,&todo.Content)
	if err!= nil {
		return 0,errors.New("error")
	}
	id,_ := result.LastInsertId()

	return id, nil
}

func (tr *TodoRepo) Get() []entity.Todo {
	result,_ := tr.DB.Query("SELECT ID,TITLE,CONTENT,IS_DONE,CREATE_AT FROM todo.todo")
	defer result.Close()
	var todolist []entity.Todo
	for result.Next(){
		var todo entity.Todo
		err := result.Scan(
			&todo.Id,
			&todo.Title,
			&todo.Content,
			&todo.IsDone,
			&todo.CreateAt,
		)
		if err != nil {
			log.Print("error")
		}
		todolist = append(todolist,todo)

	}
	return todolist

}

func (tr *TodoRepo) Del(Id int64) (entity.Todo, error) {
	getbyID,err := tr.GetById(Id)
	if err != nil{
		return getbyID,err
	}
	
	result,err := tr.DB.Query("DELETE  FROM todo.todo WHERE ID = ?" , Id)
	defer result.Close()
	if err != nil{
		return getbyID,err
	}

	return getbyID,nil
}

func changeToQueury (query map[string]interface{}) string{
	set := make([]string,0,len(query))

	for k,v := range query {
		key := stringy.New(k)
		sneakCase := key.SnakeCase("?","")
		valueType := fmt.Sprintf("%T", v)
		if valueType =="bool" {
			set = append(set,sneakCase.ToUpper()+"="+fmt.Sprintf("%v",v))
		} else {
			set = append(set,sneakCase.ToUpper()+"="+fmt.Sprintf("'%v'",v))
		}
		
	}
	return strings.Join(set,", ")
}

func (tr *TodoRepo) Update(Id int64, todo map[string]interface{}) (entity.Todo,error){
	var todoList entity.Todo
	 sql := "UPDATE todo.todo SET "+changeToQueury(todo)+" WHERE ID = ?"

	 log.Print(sql)
	 stm,err := tr.DB.Prepare(sql)	 
	 if err != nil {
		 return todoList, err
	 }
	 defer stm.Close()

	_, err = stm.Exec(Id)
	 if err != nil{
		 return todoList,errors.New("can't process update data")
	 }

	 todoList,err = tr.GetById(Id) 
	 if err != nil {
		return todoList, errors.New("can't process update data, id not found")
	}


	return todoList, nil
}

func (tr *TodoRepo) GetById(Id int64) (entity.Todo,error) {
	result := tr.DB.QueryRow("SELECT ID,TITLE,CONTENT,IS_DONE,CREATE_AT FROM todo.todo WHERE ID = ? ", Id)

	var todo entity.Todo
		err := result.Scan(
			&todo.Id,
			&todo.Title,
			&todo.Content,
			&todo.IsDone,
			&todo.CreateAt,
		)
		if err != nil {
			log.Print(todo)
			return todo,err
		}

	return todo,nil

}