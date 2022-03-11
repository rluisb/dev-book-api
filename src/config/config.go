package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//It's the connection string to database
	DatabaseConnectionString = ""

	//Port where the API will be exposed
	Port = 0

	//SecretKey is the key that will be used to sign JWT token
	SecretKey []byte
)

//This function will initialize environment variables
func Load() {
	var error error

	if error = godotenv.Load(); error != nil {
		log.Fatal(error)
	}

	Port, error = strconv.Atoi(os.Getenv("API_PORT"))
	if error != nil {
		Port = 3000
	}

	DatabaseConnectionString = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	SecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
}