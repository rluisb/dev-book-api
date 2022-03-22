package routes

import (
	"api/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		URI: "/users",
		Method: http.MethodPost,
		Function: controllers.CreateUser,
		AuthorizationRequired: false,
	},
	{
		URI: "/users",
		Method: http.MethodGet,
		Function: controllers.FindAllUsers,
		AuthorizationRequired: true,
	},
	{
		URI: "/users/{id}",
		Method: http.MethodGet,
		Function: controllers.FindUserById,
		AuthorizationRequired: true,
	},
	{
		URI: "/users/{id}",
		Method: http.MethodPut,
		Function: controllers.UpdateUser,
		AuthorizationRequired: true,
	},
	{
		URI: "/users/{id}",
		Method: http.MethodDelete,
		Function: controllers.DeleteUser,
		AuthorizationRequired: true,
	},
	{
		URI: "/users/{id}/follow",
		Method: http.MethodPost,
		Function: controllers.FollowUser,
		AuthorizationRequired: true,
	},
	{
		URI: "/users/{id}/unfollow",
		Method: http.MethodPost,
		Function: controllers.UnfollowUser,
		AuthorizationRequired: true,
	},
	{
		URI: "/users/{id}/followers",
		Method: http.MethodGet,
		Function: controllers.FindFollowers,
		AuthorizationRequired: true,
	},
}