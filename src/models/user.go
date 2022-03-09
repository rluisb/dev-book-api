package models

import (
	"errors"
	"strings"
	"time"
)

//Represents the user that will use the Social Network
type User struct {
	ID uint64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Nick string `json:"nick,omitempty"`
	Email string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

//Prepare the user by validating and formatting data to DB
func (user *User) Prepare(step string) error {
	if error := user.validate(step); error != nil {
		return error
	}

	user.format()
	return nil
}

func (user *User) validate(step string) error {
	if user.Name == "" {
		 return errors.New("User name is required and cannot be empty")
	}

	if user.Nick == "" {
		 return errors.New("User nick is required and cannot be empty")
	}

	if user.Email == "" {
		 return errors.New("User email is required and cannot be empty")
	}

	if step == "create" && user.Password == "" {
		 return errors.New("User name is required and cannot be empty")
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}

