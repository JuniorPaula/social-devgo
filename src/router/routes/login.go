package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var loginRoutes = []Route{
	{
		URI:     "/",
		Method:  http.MethodGet,
		Func:    controllers.LoginPage,
		HasAuth: false,
	},
	{
		URI:     "/login",
		Method:  http.MethodGet,
		Func:    controllers.LoginPage,
		HasAuth: false,
	},
}
