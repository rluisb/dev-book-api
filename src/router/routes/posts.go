package routes

import (
	"api/src/controllers"
	"net/http"
)

var postsRoutes = []Route {
	{
		URI: "/posts",
		Method: http.MethodPost,
		Function: controllers.CreatePost,
		AuthorizationRequired: true,
	},
	{
		URI: "/posts",
		Method: http.MethodGet,
		Function: controllers.GetPosts,
		AuthorizationRequired: true,
	},
	{
		URI: "/posts/{id}",
		Method: http.MethodGet,
		Function: controllers.GetPostById,
		AuthorizationRequired: true,
	},
	{
		URI: "/posts/{id}",
		Method: http.MethodPut,
		Function: controllers.UpdatePost,
		AuthorizationRequired: true,
	},
	{
		URI: "/posts/{id}",
		Method: http.MethodDelete,
		Function: controllers.DeletePost,
		AuthorizationRequired: true,
	},
}