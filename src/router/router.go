package router

import (
	"webapp/src/router/routes"

	"github.com/gorilla/mux"
)

func Handler() *mux.Router {
	r := mux.NewRouter()
	return routes.BootsrapRoutes(r)
}
