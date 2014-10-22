package controller

// This handles rendering our unauthenticated user facing static web pages.

import (
	"bursa.io/models"
	"bursa.io/renaissance/picasso"
	"bursa.io/renaissance/session"
	"net/http"
)

type HomeController struct{}

func (h *HomeController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Temporary command to get the ball rolling
	picasso := picasso.New(w, r)
	picasso.Render("index.html", nil)
}

// Creates a new user when they complete the signup process
func (h *HomeController) CreateUser(w http.ResponseWriter, r *http.Request) {

	// get email/password from the form
	email, pass := getCredentials(r)

	// store user in the database
	models.CreateUser(email, pass)

	session.CreateUserSession(w, r)

	// direct user to the app
	http.Redirect(w, r, "/app", http.StatusOK)
}

// validates a user's login credentials
func (h *HomeController) Login(w http.ResponseWriter, r *http.Request) {

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
