package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Create a new user on DB
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		log.Fatal(error)
	}

	var user models.User
	if error = json.Unmarshal(requestBody, &user); error != nil {
		log.Fatal(error)
	}

	db, error := database.Connect()
	if error != nil {
		log.Fatal(error)
	}

	repository := repositories.NewUsersRepository(db)

	userID, error := repository.Create(user)
	if error != nil {
		log.Fatal(error)
	}

	w.Write([]byte(fmt.Sprintf("InsertedId: %d", userID)))
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