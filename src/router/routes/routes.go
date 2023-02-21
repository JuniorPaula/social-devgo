package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI     string
	Method  string
	Func    func(http.ResponseWriter, *http.Request)
	HasAuth bool
}

func BootsrapRoutes(router *mux.Router) *mux.Router {
	routes := loginRoutes
	routes = append(routes, usersRoutes...)
	routes = append(routes, homeRoute)

	for _, route := range routes {
		router.HandleFunc(route.URI, route.Func).Methods(route.Method)
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
