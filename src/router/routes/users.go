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
	{
		URI:     "/users",
		Method:  http.MethodPost,
		Func:    controllers.CreateUser,
		HasAuth: false,
	},
	{
		URI:     "/q",
		Method:  http.MethodGet,
		Func:    controllers.FindUsersPage,
		HasAuth: true,
	},
}
