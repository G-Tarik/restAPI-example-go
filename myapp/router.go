package myapp

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		log.Println(route.Name)
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

		endpointIndex = append(endpointIndex, Endpoint{route.Name, route.Method, route.Pattern})
	}
	return router
}
