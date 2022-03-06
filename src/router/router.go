package router

import "github.com/gorilla/mux"

//Generate will return a new instance of mux Router
func Generate() *mux.Router {
	return mux.NewRouter()
}