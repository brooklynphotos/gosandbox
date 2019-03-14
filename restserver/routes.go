package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name    string
	Method  string
	Pattern string
	Handler http.HandlerFunc
}

type Routes []Route

func MakeRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, r := range routes {
		handler := Logger(r.Name, r.Handler)
		router.
			Methods(r.Method).
			Path(r.Pattern).
			Name(r.Name).
			Handler(handler)
	}
	return router
}

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
		"TodoIndexVar",
		"GET",
		"/todos/{id}",
		TodoIndexVar,
	},
}