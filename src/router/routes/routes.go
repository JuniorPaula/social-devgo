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

	for _, route := range routes {
		router.HandleFunc(route.URI, route.Func).Methods(route.Method)
	}

	return router
}
