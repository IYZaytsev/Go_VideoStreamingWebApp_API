package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//NewRouter uses the routes slice declared in routes.go and creates a router  instance
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	for _, route := range routes {
		var handler http.Handler
		handler = cors.Default().Handler(route.HandlerFunc)
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
