package controller

// This handles rendering our unauthenticated user facing static web pages.

import (
	"bursa.io/renaissance/authentication"
	"bursa.io/renaissance/models"
	"bursa.io/renaissance/picasso"
	"bursa.io/renaissance/satchel"
	"github.com/gorilla/sessions"
	"net/http"
)

type HomeController struct{}

// Returns the site home page
func (h *HomeController) Trigger(s *satchel.Satchel) {
	picasso := satchel.GetPicasso()
	picasso.Render("index.html", nil)
}

// Creates a new user when they complete the signup process
func (h *HomeController) CreateUser(s *satchel.Satchel) {

	// get email/password from the form
	writer, request := satchel.Context()
	email, pass := getCredentials(request)

	user.CreateUser(email, pass)

	// direct user to the app
	http.Redirect(writer, request, "/app", 200)
}

// validates a user's login credentials
func (h *HomeController) Login(s *satchel.Satchel) {

	// get email/password from the form
	writer, request := satchel.Context()
	email, password := getCredentials(request)

	// if not logged in successfully, return to main page
	if success := user.Authenticated(email, password); success {
		http.Redirect(writer, request, "/index.html")
	}

	// direct the user to the app if login successful
	http.Redirect(writer, request, "/app", 200)
}

// Used to fetch the a username/password from an input form
func getCredentials(r *http.Request) (string, string) {
	email := r.PostFormValues("email")
	password := r.PostFormValues("password")
	return email, password
}
