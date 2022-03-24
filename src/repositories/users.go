package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
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

//Find all users that matches the name or the nick passed
func (repository users) FindAll(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	rows, error := repository.db.Query(
		"SELECT id, name, nick, email, createdAt FROM users WHERE name like ? OR nick like ?",
		nameOrNick,
		nameOrNick,
	)
	 if error != nil {
		 return nil, error
	 }

	 defer rows.Close()

	 var users []models.User

	 for rows.Next() {
		 var user models.User

		 if error = rows.Scan(
			 &user.ID,
			 &user.Name,
			 &user.Nick,
			 &user.Email,
			 &user.CreatedAt,
		 ); error != nil {
			 return nil, error
		 }

		 users = append(users, user)
	 }

	 return users, nil
}

//Find a user by it's ID
func (repository users) FindByID(ID uint64) (models.User, error) {

	row, error := repository.db.Query(
		"SELECT id, name, nick, email, createdAt FROM users WHERE id = ? ",
		ID,
	)
	 if error != nil {
		 return models.User{}, error
	 }

	 defer row.Close()

	 var user models.User

	 if row.Next() {
		if error = row.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); error != nil {
			return models.User{}, error
		}
	 }

	 return user, nil
}

//Update an user by it's ID
func (repository users) Update(ID uint64, user models.User) error {
	statement, error := repository.db.Prepare(
		"UPDATE users SET name = ?, nick = ?, email = ? where id = ?",
	)
	if error != nil {
		 return error
	 }

	 defer statement.Close()

	 if _, error := statement.Exec(user.Name, user.Nick, user.Email, ID); error != nil {
		 return error
	 }

	 return nil
}

//Delete an user by it's ID
func (repository users) Delete(ID uint64) error {
	statement, error := repository.db.Prepare(
		"DELETE FROM users WHERE id = ?",
	)
	if error != nil {
		 return error
	 }

	 defer statement.Close()

	 if _, error := statement.Exec(ID); error != nil {
		 return error
	 }

	 return nil
}

//Find a user by email and return it's id and hashed password
func (repository users) FindUserByEmail(email string) (models.User, error) {
	row, error := repository.db.Query("SELECT id, password FROM users WHERE email = ?", email)
	if error != nil {
		return models.User{}, error
	}
	defer row.Close()

	var user models.User

	if row.Next() {
		if error = row.Scan(&user.ID, &user.Password); error != nil {
			return models.User{}, error
		}
	}

	return user, nil
}

//Create following between users
func (repository users) Follow(userID, followerID uint64) error {
	statement, error := repository.db.Prepare(
		"INSERT IGNORE INTO followers (user_id, follower_id) VALUES (?, ?)",
	)

	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error = statement.Exec(userID, followerID); error != nil {
		return error
	}

	return nil
}

//Remove following between users
func (repository users) Unfollow(userID, followerID uint64) error {
	statement, error := repository.db.Prepare(
		"DELETE FROM followers WHERE user_id = ? AND follower_id = ?",
	)

	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error = statement.Exec(userID, followerID); error != nil {
		return error
	}

	return nil
}

//Find followers
func (repository users) FindFollowersByUserId(userID uint64) ([]models.User, error) {
	rows, error := repository.db.Query(`
		SELECT u.id, u.name, u.nick, u.email, u.createdAt
		FROM users u
		INNER JOIN followers f ON u.id = f.follower_id
		WHERE f.user_id = ?
	`, userID)

	if error != nil {
		 return nil, error
	 }

	defer rows.Close()

	 var users []models.User

	 for rows.Next() {
		 var user models.User

		 if error = rows.Scan(
			 &user.ID,
			 &user.Name,
			 &user.Nick,
			 &user.Email,
			 &user.CreatedAt,
		 ); error != nil {
			 return nil, error
		 }

		 users = append(users, user)
	 }

	 return users, nil
}

//Find following
func (repository users) FindFollowingByUserId(userID uint64) ([]models.User, error) {
	rows, error := repository.db.Query(`
		SELECT u.id, u.name, u.nick, u.email, u.createdAt
		FROM users u
		INNER JOIN followers f ON u.id = f.user_id
		WHERE f.follower_id = ?
	`, userID)

	if error != nil {
		 return nil, error
	 }

	defer rows.Close()

	 var users []models.User

	 for rows.Next() {
		 var user models.User

		 if error = rows.Scan(
			 &user.ID,
			 &user.Name,
			 &user.Nick,
			 &user.Email,
			 &user.CreatedAt,
		 ); error != nil {
			 return nil, error
		 }

		 users = append(users, user)
	 }

	 return users, nil
}

//Find user's password by userID
func (repository users) FindPasswordByUserId(userID uint64) (string, error) {
	row, error := repository.db.Query("SELECT password FROM users WHERE id = ?", userID)
	if error != nil {
		return "", error
	}
	defer row.Close()

	var user models.User

	if row.Next() {
		if error = row.Scan(&user.Password); error != nil {
			return "", error
		}
	}

	return user.Password, nil
}

//Update user's password
func (repository users) UpdatePassword(userID uint64, password string) error {
	statement, error := repository.db.Prepare("UPDATE users SET password = ? WHERE id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error = statement.Exec(password, userID); error != nil {
		return error
	}

	return nil
}