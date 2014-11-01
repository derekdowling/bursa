package home

// This handles rendering our unauthenticated user facing static web pages.

import (
	"bursa.io/email"
	"bursa.io/models"
	"bursa.io/picasso"
	"bursa.io/renaissance/session"
	"net/http"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	// Temporary command to get the ball rolling
	picasso.Render(w, "marketing/layout", "marketing/index", nil)
}

func HandleSignup(w http.ResponseWriter, r *http.Request) {
	email := r.PostFormValue("email")
	if email == "" {
		// TODO: make this redirect back to the signup page
	}
	return email.Subscribe(email)
}

// Creates a new user when they complete the signup process
func HandleCreateUser(w http.ResponseWriter, r *http.Request) {

	// get email/password from the form
	email, pass := getCredentials(r)

	// store user in the database
	models.CreateUser(email, pass)

	session.CreateUserSession(w, r)

	// direct user to the app
	http.Redirect(w, r, "/app", http.StatusOK)
}

// validates a user's login credentials
func HandleLogin(w http.ResponseWriter, r *http.Request) {

	// get email/password from the form
	email, password := getCredentials(r)

	// if not logged in successfully, return to main page
	user := models.AttemptLogin(email, password)

	if user == nil {
		// TODO: include attempted email for auto-fill
		// TODO: add login fail flag for login failure alert
		http.Redirect(w, r, "/index.html", http.StatusUnauthorized)
	}

	// direct the user to the app if login successful
	http.Redirect(w, r, "/app", http.StatusOK)
}

// Used to fetch the a username/password from an input form
func getCredentials(r *http.Request) (string, string) {
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	return email, password
}
