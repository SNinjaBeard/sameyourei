package server

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"strconv"

	"github.com/SuperSpaceNinja/sameyourei/model"
	"github.com/SuperSpaceNinja/sameyourei/repository"
	"github.com/gorilla/mux"
)

const contentType string = "Content-Type"
const mimetypeJSONUtf8 string = "application/json; charset=UTF-8"

/*
Index Welcome */
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handlers.Index()")
	fmt.Fprintln(w, "Welcome!")
}

/*
TodoIndex get all todos */
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handlers.TodoIndex()")
	w.Header().Set(contentType, mimetypeJSONUtf8)

	var todos model.Todos
	todos, err := repository.RepoGetAllTodos()
	if err != nil {
		respe(w, r, http.StatusInternalServerError, err)
		return
	}

	respt(w, r, http.StatusOK, todos)
}

/*
TodoByID get todo by id */
func TodoByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handlers.TodoByID()")
	w.Header().Set(contentType, mimetypeJSONUtf8)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["todoID"])
	if err != nil {
		respe(w, r, http.StatusBadRequest, err)
		return
	}

	todos, err := repository.RepoFindTodoByID(id)
	if err != nil {
		respe(w, r, http.StatusInternalServerError, err)
		return
	}

	respt(w, r, http.StatusOK, todos)
}

/*
TodoCreate creates todo */
func TodoCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handlers.TodoCreate()")
	w.Header().Set(contentType, mimetypeJSONUtf8)

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	defer r.Body.Close()
	if err != nil {
		respe(w, r, http.StatusBadRequest, err)
		return
	}

	var todo model.Todo
	if err := json.Unmarshal(body, &todo); err != nil {
		respe(w, r, http.StatusUnprocessableEntity, err)
		return
	}

	todos, err := repository.RepoCreateTodo(todo)
	if err != nil {
		respe(w, r, http.StatusInternalServerError, err)
		return
	}

	respt(w, r, http.StatusCreated, todos)
}

/*
TodoUpdate updates todo */
func TodoUpdate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handlers.TodoUpdate()")
	w.Header().Set(contentType, mimetypeJSONUtf8)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["todoID"])
	if err != nil {
		respe(w, r, http.StatusBadRequest, err)
		return
	}

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	defer r.Body.Close()
	if err != nil {
		respe(w, r, http.StatusBadRequest, err)
		return
	}

	var todo model.Todo
	if err := json.Unmarshal(body, &todo); err != nil {
		respe(w, r, http.StatusUnprocessableEntity, err)
		return
	}

	todos, err := repository.RepoUpdateTodo(id, todo)
	if err != nil {
		respe(w, r, http.StatusInternalServerError, err)
		return
	}

	respt(w, r, http.StatusOK, todos)
}

/*
TodoDelete remove todo */
func TodoDelete(w http.ResponseWriter, r *http.Request) {

}
