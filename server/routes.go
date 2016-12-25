package server

import "net/http"

/*
Route route struct */
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

/*
Routes route array */
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		TodoIndex,
	},
	Route{
		"TodoById",
		"GET",
		"/todos/{todoID}",
		TodoByID,
	},
	Route{
		"TodoCreate",
		"POST",
		"/todos",
		TodoCreate,
	},
	Route{
		"TodoDelete",
		"DELETE",
		"/todos/{todoID}",
		TodoDelete,
	},
	Route{
		"TodoUpdate",
		"PUT",
		"/todos/{todoID}",
		TodoUpdate,
	},
}
