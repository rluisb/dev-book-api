package routes

import (
	"api/src/controller"
	"net/http"
)

var usersRoutes = []Route{
	{
		URI: "/users",
		Method: http.MethodPost,
		Function: controller.CreateUser,
		AuthorizationRequired: false,
	},
	{
		URI: "/users",
		Method: http.MethodGet,
		Function: controller.FindAllUsers,
		AuthorizationRequired: false,
	},
	{
		URI: "/users/{id}",
		Method: http.MethodGet,
		Function: controller.FindUserById,
		AuthorizationRequired: false,
	},
	{
		URI: "/users/{id}",
		Method: http.MethodPut,
		Function: controller.UpdateUser,
		AuthorizationRequired: false,
	},
	{
		URI: "/users/{id}",
		Method: http.MethodDelete,
		Function: controller.DeleteUser,
		AuthorizationRequired: false,
	},
}