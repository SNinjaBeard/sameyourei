package repository

import (
	"database/sql"
	"fmt"

	"log"

	"time"

	"github.com/SuperSpaceNinja/sameyourei/db"
	"github.com/SuperSpaceNinja/sameyourei/model"
	_ "github.com/lib/pq"
)

/*
init opens db connectionpool and returns an sql.DB struct to
manage connectionpool */
func init() {
	fmt.Println("repository.init()")

	var err error
	db.Conn, err = sql.Open("postgres", "user=postgres dbname=homefkndb sslmode=disable")
	if err != nil {
		log.Fatal("Error: data source arguments not valid")
	}
	fmt.Println("login successful")

	err = db.Conn.Ping()
	if err != nil {
		log.Fatal("Error: could not establish connection")
	}
	fmt.Println("connection successful")
}

/*
RepoGetAllTodos get all saved todos */
func RepoGetAllTodos() (model.Todos, error) {
	fmt.Println("repository.RepoGetAllTodos()")

	var todos model.Todos
	rows, err := db.Conn.Query(
		"SELECT id, name, completed, due, created FROM sameyourei.todo")
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
		return todos, err
	}

	var todo model.Todo
	for rows.Next() {
		err := rows.Scan(&todo.ID, &todo.Name, &todo.Completed, &todo.Due, &todo.Created)
		if err != nil {
			log.Fatal(err)
			return todos, err
		}

		todos = append(todos, todo)
	}

	return todos, err
}

/*
RepoFindTodoByID find todo by id */
func RepoFindTodoByID(id int) (model.Todos, error) {
	fmt.Println("repository.RepoFindTodoByID()")

	var todos model.Todos
	rows, err := db.Conn.Query(
		"SELECT id, name, completed, due, created FROM sameyourei.todo WHERE id = $1",
		id)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
		return todos, err
	}

	var todo model.Todo
	for rows.Next() {
		err := rows.Scan(&todo.ID, &todo.Name, &todo.Completed, &todo.Due, &todo.Created)
		if err != nil {
			log.Fatal(err)
			return todos, err
		}

		todos = append(todos, todo)
	}

	if todo.ID <= 0 {
		err = sql.ErrNoRows
	}

	return todos, err
}

/*
RepoCreateTodo create todo */
func RepoCreateTodo(todo model.Todo) (model.Todos, error) {
	fmt.Println("repository.RepoCreateTodo()")

	var todos model.Todos
	rows, err := db.Conn.Query(
		"INSERT INTO sameyourei.todo(name, completed, due, created) VALUES($1, $2, $3, $4) RETURNING id, name, completed, due, created",
		todo.Name, false, todo.Due, time.Now())
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
		return todos, err
	}

	for rows.Next() {
		err = rows.Scan(&todo.ID, &todo.Name, &todo.Completed, &todo.Due, &todo.Created)
		if err != nil {
			log.Fatal(err)
			return todos, err
		}

		todos = append(todos, todo)
	}

	if todo.ID <= 0 {
		err = sql.ErrNoRows
	}

	return todos, err
}

/*
RepoUpdateTodo updates todo */
func RepoUpdateTodo(id int, todo model.Todo) (model.Todos, error) {
	fmt.Println("repository.RepoUpdateTodo()")

	// TODO: if no value changed keep previous value

	var name string

	if todo.Name != "" {
		name = todo.Name
	}

	fmt.Printf("todo.Completed: %t\n", todo.Completed)

	var todos model.Todos
	rows, err := db.Conn.Query(
		"UPDATE sameyourei.todo SET name = $2, completed = $3, due = $4 WHERE id = $1 RETURNING id, name, completed, due, created",
		id, name, todo.Completed, todo.Due)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
		return todos, err
	}

	for rows.Next() {
		err = rows.Scan(&todo.ID, &todo.Name, &todo.Completed, &todo.Due, &todo.Created)
		if err != nil {
			log.Fatal(err)
			return todos, err
		}

		todos = append(todos, todo)
	}

	if todo.ID <= 0 {
		err = sql.ErrNoRows
	}

	return todos, err
}

/*
RepoDestroyTodo remove todo */
func RepoDestroyTodo(id int) (model.Todos, error) {
	fmt.Println("repository.RepoDestroyTodo()")
	var todos model.Todos

	return todos, nil
}
