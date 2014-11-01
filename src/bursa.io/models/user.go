package models

// Handles all things user related such as creating them, logging in, updating
// attributes, and deleting them

import (
	"bursa.io/config"
	"bursa.io/renaissance/authentication"
	"github.com/mattbaird/gochimp"
	"log"
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
	db, _ := Connect()
	db.Create(&user)
}

// Test's whether or not a user has authenticated successfully
func AttemptLogin(email string, password string) *User {
	// Todo: this is broken
	db, _ := Connect()
	var user *User
	db.Where("email = ?", email).First(&user)
	// match := authentication.PasswordMatch(password, user.salt, user.hash)
	return user
}

// Adds a user, via their email, to one of our MailChimp mailing lists
func SubscribeToMail(userEmail string) gochimp.Email {
	chimp := getMailChimp()
	request := gochimp.ListsSubscribe{
		ListId:         getMailListId(),
		Email:          gochimp.Email{Email: userEmail},
		DoubleOptIn:    false,
		UpdateExisting: true,
		SendWelcome:    sendWelcomeEmail(),
	}

	resp, err := chimp.ListsSubscribe(request)
	if err != nil {
		log.Println(err.Error())
		return gochimp.Email{}
	}
	return resp
}

// Checks whether or not we are in production to avoid spamming ourselves
// with email
func sendWelcomeEmail() bool {
	if config.IsSet("production") {
		return true
	}
	return false
}

// Sets up the MailChimp API
func getMailChimp() *gochimp.ChimpAPI {
	api_key := config.GetStringMapString("email")["mailchimp_key"]
	return gochimp.NewChimp(api_key, true)
}

// Determines which mailing list to add user to based on context
func getMailListId() string {
	if config.IsSet("production") {
		return config.GetStringMapString("mail_list")["production"]
	}
	return config.GetStringMapString("mail_list")["dev"]
}
