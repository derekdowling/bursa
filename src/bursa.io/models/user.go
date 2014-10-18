package models

// Handles all things user related such as creating them, logging in, updating
// attributes, and deleting them

import (
    "github.com/gorilla/sessions"
	"bursa.io/models"
)

session.Options = &sessions.Options{
	Path:     "/",
	MaxAge:   86400 * 7,
	HttpOnly: true,
}

var store = sessions.NewCookieStore([]byte(viper.Get("session-key")))

func CreateUser(email string, password string) {

	// hash and salt password
	password_obj := authentication.CreatePassword(password)

	// create user object
	user := User{
		email:    email,
		password: password_obj,
	}

	// create/save user
	db.Create(&user)
}

// Test's whether or not a user has authenticated successfully
func Authenticated(email string, password string) bool {
	user := db.Where("email = ?", email)
	match := authentication.PasswordMatch(password, user.password)

	return match
}

func createUserSession() {
}
