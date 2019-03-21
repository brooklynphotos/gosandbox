package main

import "github.com/gorilla/mux"

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
