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

func SubscribeToMail(userEmail string) gochimp.Email {
	chimp := getMailChimp()
	request := gochimp.ListsSubscribe{
		ListId:         getMailListId(),
		Email:          gochimp.Email{Email: userEmail},
		DoubleOptIn:    false,
		UpdateExisting: true,
	}

	resp, err := chimp.ListsSubscribe(request)
	if err != nil {
		log.Println(err.Error())
		return gochimp.Email{}
	}
	return resp
}

func getMailChimp() *gochimp.ChimpAPI {
	api_key := config.GetStringMapString("email")["mailchimp_key"]
	return gochimp.NewChimp(api_key, true)
}

func getMailListId() string {
	if config.IsSet("production") {
		return config.GetStringMapString("mail_list")["production"]
	}
	return config.GetStringMapString("mail_list")["dev"]
}
