package controller

// This handles rendering our unauthenticated user facing static web pages.

import (
	"bursa.io/models"
	"bursa.io/renaissance/authentication"
	"bursa.io/renaissance/picasso"
	"net/http"
)

type HomeController struct{}

func (h *HomeController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Temporary command to get the ball rolling
	picasso := picasso.New(w, r)
	picasso.Render("index.html", nil)
}

// Creates a new user when they complete the signup process
func (h *HomeController) CreateUser(s *satchel.Satchel) {

	// get email/password from the form
	writer, request := satchel.Context()
	email, pass := getCredentials(request)

	// store user in the database
	user.CreateUser(email, pass)

	models.CreateUserSession()

	// direct user to the app
	http.Redirect(writer, request, "/app", 200)
}

// validates a user's login credentials
func (h *HomeController) Login(s *satchel.Satchel) {

	// get email/password from the form
	writer, request := satchel.Context()
	email, password := getCredentials(request)

	// if not logged in successfully, return to main page
	if user := user.AttemptLogin(email, password); user {
		// TODO: include attempted email for auto-fill
		// TODO: add login fail flag for login failure alert
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
