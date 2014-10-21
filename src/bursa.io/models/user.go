package models

// Handles all things user related such as creating them, logging in, updating
// attributes, and deleting them

import (
	"bursa.io/renaissance/authentication"
)

type Role int

// Our role definitions as specified by renaissance/firewall
const (
	Visitor Role = 1 << iota
	Authenticated
)

func CreateUser(email string, password string) {

	// hash and salt password
	salt, hash := authentication.CreatePassword(password)

	// create user object
	user := &User{
		Email:    email,
		Salt:     salt,
		Password: hash,
	}

	// create/save user
	db := Connect()
	db.Create(&user)
}

// Test's whether or not a user has authenticated successfully
func AttemptLogin(email string, password string) User {
	// Todo: this is broken
	db := Connect()
	user := db.Where("email = ?", email)
	match := authentication.PasswordMatch(password, user.salt, user.hash)
	return user
}
