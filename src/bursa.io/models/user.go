package models

// Handles all things user related such as creating them, logging in, updating
// attributes, and deleting them

import (
	"bursa.io/renaissance/authentication"
)

func CreateUser(email string, password string) {

	// hash and salt password
	password_obj := authentication.CreatePassword(password)

	// create user object
	user := User{
		email:    email,
		password: password_obj,
	}

	// create/save user
	db := models.Connect()
	db.Create(&user)
}

// Test's whether or not a user has authenticated successfully
func ValidCredentials(email string, password string) bool {
	db := models.Connect()
	user := db.Where("email = ?", email)
	match := authentication.PasswordMatch(password, user.password)
	return match
}
