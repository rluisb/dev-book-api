package repositories

import (
	"api/src/models"
	"database/sql"
)

//Represents an users repository
type users struct {
	db *sql.DB
}

//This function creates an users repository by using something like a Dependency Injection
func NewUsersRepository(db *sql.DB) *users {
	return &users{db}
}

//Create a new user on DB
func (repository users) Create (user models.User) (uint64, error) {
	 statement, error := repository.db.Prepare("INSERT INTO users (name, nick, email, password) VALUES (?, ?, ?, ?)")
	 if error != nil {
		 return 0, error
	 }

	 defer statement.Close()

	 resultSet, error := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	 if error != nil {
		 return 0, error
	 }

	 lastInsertID, error := resultSet.LastInsertId()
	 if error != nil {
		 return 0, error
	 }

	 return uint64(lastInsertID), nil
}