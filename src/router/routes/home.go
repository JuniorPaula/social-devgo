package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var homeRoute = Route{
	URI:     "/home",
	Method:  http.MethodGet,
	Func:    controllers.HomePage,
	HasAuth: true,
}
