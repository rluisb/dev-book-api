package controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//Create a new user on DB
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating an User"))
}

//Find all users on DB
func FindAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Finding all users"))
}

//Find a specific user on DB based on it's Id
func FindUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, ok := vars["id"]

	if !ok {
        fmt.Println("id is missing in parameters")
    }

	w.Write([]byte(fmt.Sprintf("Find user by ID: %s", id)))
}

//Update user information on DB
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, ok := vars["id"]

	if !ok {
        fmt.Println("id is missing in parameters")
    }

	w.Write([]byte(fmt.Sprintf("Updating an User by ID: %s", id)))
}

//Delete user information from DB
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, ok := vars["id"]

	if !ok {
        fmt.Println("id is missing in parameters")
    }

	w.Write([]byte(fmt.Sprintf("Deleting an User by ID: %s", id)))
}