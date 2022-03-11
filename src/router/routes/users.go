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
		AuthorizationRequired: false,
	},
	{
		URI: "/users/{id}",
		Method: http.MethodPut,
		Function: controllers.UpdateUser,
		AuthorizationRequired: false,
	},
	{
		URI: "/users/{id}",
		Method: http.MethodDelete,
		Function: controllers.DeleteUser,
		AuthorizationRequired: false,
	},
}