package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

//Route represents all routes structure from the API
type Route struct {
	URI string
	Method string
	Function func(http.ResponseWriter, *http.Request)
	AuthorizationRequired bool
}

//Configure router by adding all routes inside mux.Router
func Configure(router *mux.Router) *mux.Router {
	routes := usersRoutes
	routes = append(routes, loginRoute)
	routes = append(routes, postsRoutes...)

	for _, route := range routes {
		if route.AuthorizationRequired {
			router.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Authenticate(route.Function)),
			).Methods(route.Method)
		} else {
			router.HandleFunc(route.URI,
				middlewares.Logger(route.Function),
			).Methods(route.Method)
		}
	}

	return router
}
