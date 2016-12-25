package server

import (
	"fmt"
	"net/http"

	"github.com/SuperSpaceNinja/sameyourei/util"
	"github.com/gorilla/mux"
)

/*
NewRouter constructor */
func NewRouter() *mux.Router {
	fmt.Println("server.NewRouter()")

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = util.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
