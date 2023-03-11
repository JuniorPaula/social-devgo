package routes

import (
	"net/http"
	"webapp/src/middlewares"

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
	routes = append(routes, postRoutes...)
	routes = append(routes, logoutRoute)

	for _, route := range routes {

		if route.HasAuth {
			router.HandleFunc(route.URI, middlewares.Logger(middlewares.Authenticate(route.Func))).Methods(route.Method)

		} else {
			router.HandleFunc(route.URI, middlewares.Logger(route.Func)).Methods(route.Method)
		}

	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
