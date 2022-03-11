package middlewares

import (
	"api/src/authentication"
	"api/src/responses"
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
		if error := authentication.ValidateToken(r); error != nil {
			responses.Error(w, http.StatusUnauthorized, error)
			return
		}
		next(w, r)
	}
}