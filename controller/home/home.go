// This handles rendering all of our unauthenticated user facing static web pages
package home

import (
	"github.com/derekdowling/bursa/email"
	"github.com/derekdowling/bursa/models"
	"github.com/derekdowling/bursa/picasso"
	"github.com/derekdowling/bursa/renaissance/session"
	"net/http"
)

type Form struct {
	Email string
	Name  string
}

// Handles loading the main page of the website
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	picasso.Render(w, "marketing/layout", "marketing/index", nil)
}

// Completes a user signup. Assumes that the values being provided from the
// front-end have already been validated
func HandleAbout(w http.ResponseWriter, r *http.Request) {
	picasso.Render(w, "marketing/layout", "marketing/about", struct{}{})
}

func HandleSignup(w http.ResponseWriter, r *http.Request) {

	userEmail := r.FormValue("email")
	success := email.Subscribe(userEmail)

	if !success {
		picasso.RenderWithCode(w, "marketing/layout", "marketing/index", userEmail, http.StatusBadRequest)
	}

	picasso.Render(w, "marketing/layout", "marketing/success", nil)
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

func HandleSignupSuccess(w http.ResponseWriter, r *http.Request) {
	picasso.Render(w, "marketing/layout", "marketing/success", nil)
}

func Handle404(w http.ResponseWriter, r *http.Request) {
	picasso.Render(w, "marketing/layout", "marketing/404", nil)
}

// Used to fetch the a username/password from an input form
func getCredentials(r *http.Request) (string, string) {
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	return email, password
}