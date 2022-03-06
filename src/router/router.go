package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

//Generate will return a new instance of mux Router
func Generate() *mux.Router {
	router := mux.NewRouter()
	return routes.Configure(router)
}