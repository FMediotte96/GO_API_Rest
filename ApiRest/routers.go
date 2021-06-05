package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			HandlerFunc(route.HandleFunc)
	}

	return router
}

var routes = Routes{
	Route{"Index", "GET", "/", Index},
	Route{"MovieList", "GET", "/peliculas", MovieList},
	Route{"MovieShow", "GET", "/peliculas/{id}", MovieShow},
	Route{"MovieAdd", "POST", "/peliculas", MovieAdd},
}
