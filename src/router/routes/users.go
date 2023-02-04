package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var usersRoutes = []Route{
	{
		URI:     "/signup-user",
		Method:  http.MethodGet,
		Func:    controllers.SignupPage,
		HasAuth: false,
	},
}