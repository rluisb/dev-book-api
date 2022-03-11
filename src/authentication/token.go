package authentication

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//Generates a token with user permissions
func GenerateToken(userID uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return token.SignedString(config.SecretKey)
}

//Check if token in request is valid
func ValidateToken(r *http.Request) error {
	tokenFromRequest := getTokenFromRequest(r)
	token, error := jwt.Parse(tokenFromRequest, returnVerificationKey)
	if error != nil {
		return error
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("invalid token")
}

func GetUserIDFromToken(r *http.Request) (uint64, error) {
	tokenFromRequest := getTokenFromRequest(r)
	token, error := jwt.Parse(tokenFromRequest, returnVerificationKey)
	if error != nil {
		return 0, error
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, error := strconv.ParseUint(fmt.Sprintf("%0.f", permissions["userId"]), 10, 64)
		if error != nil {
			return 0, error
		}

		return userID, nil
	}

	return 0, errors.New("invalid token")
}

func getTokenFromRequest(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) != 2 {
		return ""
	}

	return strings.Split(token, " ")[1]
}

func returnVerificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected sining method! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}