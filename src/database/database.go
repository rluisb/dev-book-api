package database

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver
)

//Open connection to DB and return it
func Connect() (*sql.DB, error) {
	db, error := sql.Open("mysql", config.DatabaseConnectionString)
	if error != nil {
		return nil, error
	}

	if error = db.Ping(); error != nil {
		db.Close()
		return nil, error
	}

	return db, nil
}