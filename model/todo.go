package model

import (
	"time"
)

type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
	Created   time.Time `json:"created"`
}

type Todos []Todo

/*
CreateTodo assembles a todo from parameters */
func CreateTodo(id int, name string, completed bool, due time.Time, created time.Time) Todo {
	var todo Todo
	todo.ID = id
	todo.Name = name
	todo.Completed = completed
	todo.Due = due
	todo.Created = created
	return todo
}
