package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

//Log requests made to server
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n[Router] %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

//Check if user that is doing the request is authenticated
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Verifying authentication")
		next(w, r)
	}
}