package models

// Handles all things user related such as creating them, logging in, updating
// attributes, and deleting them

import (
	"github.com/derekdowling/bursa/emailer"
	"github.com/derekdowling/bursa/renaissance/authentication"
	"time"
)

type Role int

// Our role definitions as specified by renaissance/firewall
const (
	Visitor Role = 1 << iota
	Authenticated
)

type User struct {
	Id        int64
	Name      string `sql:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Password  string `sql:"size:255"`
	Salt      string `sql:"size:64"`

	Email string `sql:"size:255"`
}

func CreateUser(email string, password string) int64 {

	// hash and salt password
	salt, hash := authentication.CreatePassword(password)

	// create user object
	user := &User{
		Email:    email,
		Salt:     salt,
		Password: hash,
	}

	// create user
	db := Connect()
	db.Create(&user)

	if emailer.Enabled() {
		emailer.Subscribe(email)
	}

	return user.Id
}

func FindUser(id int64) *User {
	var user *User
	db := Connect()
	db.First(&user, id)
	return user
}

func FindUserByEmail(email string) *User {
	var user *User
	db := Connect()
	db.Where("email=?", email).First(&user)
	return user
}

// Test's whether or not a user has authenticated successfully
func FindUserByCreds(email string, password string) *User {
	user := FindUserByEmail(email)

	if match := authentication.PasswordMatch(password, user.Salt, user.Password); !match {
		return nil
	}

	return user
}
