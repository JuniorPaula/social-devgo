package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var postRoutes = []Route{
	{
		URI:     "/posts",
		Method:  http.MethodPost,
		Func:    controllers.CreatePost,
		HasAuth: true,
	},
}
