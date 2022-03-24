package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//Create a post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	userIDFromToken, error := authentication.GetUserIDFromToken(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		responses.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var post models.Post
	if error = json.Unmarshal(requestBody, &post); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	post.AuthorID = userIDFromToken

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewPostsRepository(db)

	post.ID, error = repository.Create(post)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}


	responses.JSON(w, http.StatusCreated, post)
}

//List all post
func GetPosts(w http.ResponseWriter, r *http.Request) {}

//Get a post
func GetPostById(w http.ResponseWriter, r *http.Request) {}

//Update a post
func UpdatePost(w http.ResponseWriter, r *http.Request) {}

//Delete a post
func DeletePost(w http.ResponseWriter, r *http.Request) {}



