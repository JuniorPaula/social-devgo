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
	{
		URI:     "/posts/{postId}/like",
		Method:  http.MethodPost,
		Func:    controllers.LikePost,
		HasAuth: true,
	},
}
