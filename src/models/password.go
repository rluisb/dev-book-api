package models

//Represents a password
type Password struct {
	NewPassword string `json:"newPassword"`
	OldPassword string `json:"oldPassword"`
}