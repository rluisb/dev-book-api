package repositories

import (
	"api/src/models"
	"database/sql"
)

//Represents an users repository
type posts struct {
	db *sql.DB
}

//This function creates an users repository by using something like a Dependency Injection
func NewPostsRepository(db *sql.DB) *posts {
	return &posts{db}
}

//Create a new post on DB
func (repository posts) Create(post models.Post) (uint64, error) {
	 statement, error := repository.db.Prepare("INSERT INTO posts (title, content, author_id) VALUES (?, ?, ?)")
	 if error != nil {
		 return 0, error
	 }

	 defer statement.Close()

	 resultSet, error := statement.Exec(post.Title, post.Content, post.AuthorID)
	 if error != nil {
		 return 0, error
	 }

	 lastInsertID, error := resultSet.LastInsertId()
	 if error != nil {
		 return 0, error
	 }

	 return uint64(lastInsertID), nil
}

