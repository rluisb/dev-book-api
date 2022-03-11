package security

import "golang.org/x/crypto/bcrypt"

//Receives an string an turn it into a hash
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

//Check password from user by comparing the password that user sent in a request to a hashed password that exists in database
func ValidatePassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}