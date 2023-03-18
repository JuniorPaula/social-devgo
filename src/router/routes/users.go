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
	{
		URI:     "/users/{userId}",
		Method:  http.MethodGet,
		Func:    controllers.UserProfilePage,
		HasAuth: true,
	},
	{
		URI:     "/users/{userId}/unfollower",
		Method:  http.MethodPost,
		Func:    controllers.UnfollowerUser,
		HasAuth: true,
	},
	{
		URI:     "/users/{userId}/follower",
		Method:  http.MethodPost,
		Func:    controllers.FollowerUser,
		HasAuth: true,
	},
	{
		URI:     "/profile",
		Method:  http.MethodGet,
		Func:    controllers.UserLoggedProfilePage,
		HasAuth: true,
	},
	{
		URI:     "/edit-user",
		Method:  http.MethodGet,
		Func:    controllers.EditUserPage,
		HasAuth: true,
	},
	{
		URI:     "/edit-user",
		Method:  http.MethodPut,
		Func:    controllers.UpdateUser,
		HasAuth: true,
	},
}
