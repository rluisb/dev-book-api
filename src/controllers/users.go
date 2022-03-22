package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

//Create a new user on DB
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		responses.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var user models.User
	if error = json.Unmarshal(requestBody, &user); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if error = user.Prepare("create"); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)

	user.ID, error = repository.Create(user)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

//Find all users on DB
func FindAllUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := getQueryParamByName(r, "user")

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)

	users, error := repository.FindAll(nameOrNick)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

//Find a specific user on DB based on it's Id
func FindUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userID, error := strconv.ParseUint(vars["id"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)

	user, error := repository.FindByID(userID)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

//Update user information on DB
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userID, error := strconv.ParseUint(vars["id"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	userIDFromToken, error := authentication.GetUserIDFromToken(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	if userIDFromToken != userID {
		responses.Error(w, http.StatusForbidden, errors.New("it's not possible update an user that is not yours"))
		return
	}

	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		responses.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var user models.User
	if error = json.Unmarshal(requestBody, &user); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if error = user.Prepare("update"); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)

	if error = repository.Update(userID, user); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

//Delete user information from DB
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userID, error := strconv.ParseUint(vars["id"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	userIDFromToken, error := authentication.GetUserIDFromToken(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	if userIDFromToken != userID {
		responses.Error(w, http.StatusForbidden, errors.New("it's not possible delete an user that is not yours"))
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)

	if error = repository.Delete(userID); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

//Ensure that a user can follow another
func FollowUser(w http.ResponseWriter, r *http.Request) {
	followerId, error := authentication.GetUserIDFromToken(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	vars := mux.Vars(r)
	userId, error := strconv.ParseUint(vars["id"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if followerId == userId {
		responses.Error(w, http.StatusForbidden, errors.New("it's not possible to follow yourself"))
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	if error = repository.Follow(userId, followerId); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

//Ensure that a user can unfollow another
func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	followerId, error := authentication.GetUserIDFromToken(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	vars := mux.Vars(r)
	userId, error := strconv.ParseUint(vars["id"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if followerId == userId {
		responses.Error(w, http.StatusForbidden, errors.New("it's not possible to ollow yourself"))
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	if error = repository.Unfollow(userId, followerId); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

//Find all followers from a user
func FindFollowers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userID, error := strconv.ParseUint(vars["id"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)

	users, error := repository.FindFollowersByUserId(userID)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

func getQueryParamByName(r *http.Request, queryParamName string) string{
	return strings.ToLower(r.URL.Query().Get(queryParamName))
}