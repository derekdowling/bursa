// This handles rendering all of our unauthenticated user facing static web pages
package home

import (
	log "github.com/Sirupsen/logrus"
	"github.com/derekdowling/bursa/models"
	"github.com/derekdowling/bursa/picasso"
	"github.com/derekdowling/bursa/renaissance/session"
	"github.com/gorilla/schema"
	"net/http"
)

// Handles loading the main page of the website
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	picasso.Render(w, "marketing/layout", "marketing/index", nil)
}

// Completes a user signup. Assumes that the values being provided from the
// front-end have already been validated
func HandleAbout(w http.ResponseWriter, r *http.Request) {
	picasso.Render(w, "marketing/layout", "marketing/about", struct{}{})
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	picasso.Render(w, "marketing/layout", "marketing/login", struct{}{})
}

func HandleSignup(w http.ResponseWriter, r *http.Request) {
	picasso.Render(w, "marketing/layout", "marketing/signup", struct{}{})
}

func HandleForgotPassword(w http.ResponseWriter, r *http.Request) {
	picasso.Render(w, "marketing/layout", "marketing/forgot-password", struct{}{})
}

// Struct to marshall signup data into
type Credentials struct {
	Email    string
	Password string
	Company  string
}

// Marshals the results of a Post Form into our Credentials Object for either
func marshalForm(r *http.Request) *Credentials {
	// parse out our query vars
	if err := r.ParseForm(); err != nil {
		log.Error(err)
		return nil
	}

	// marshal over to a struct
	decoder := schema.NewDecoder()
	credentials := new(Credentials)

	if err := decoder.Decode(credentials, r.PostForm); err != nil {
		log.Error(err)
		return nil
	}

	return credentials
}

func HandlePostSignup(w http.ResponseWriter, r *http.Request) {

	creds := marshalForm(r)

	// store user in the database
	id := models.CreateUser(creds.Email, creds.Password)
	if id == 0 {
		picasso.RenderWithCode(
			w, "marketing/layout", "marketing/signup", creds.Email, http.StatusBadRequest,
		)
		return
	}

	// create a session and direct user to the app
	session.CreateUserSession(w, r, id)
	http.Redirect(w, r, "/app", http.StatusOK)
}

// validates a user's login credentials
func HandlePostLogin(w http.ResponseWriter, r *http.Request) {

	creds := marshalForm(r)

	// if not logged in successfully, return to main page
	user := models.FindUserByCreds(creds.Email, creds.Password)
	if user == nil {
		http.Redirect(w, r, "/login", http.StatusUnauthorized)
		return
	}

	// direct the user to the app if login successful
	session.CreateUserSession(w, r, user.Id)
	http.Redirect(w, r, "/app", http.StatusOK)
}

func Handle404(w http.ResponseWriter, r *http.Request) {
	picasso.Render(w, "marketing/layout", "marketing/404", nil)
}

func Handle500(w http.ResponseWriter, r *http.Request) {
	picasso.Render(w, "marketing/layout", "marketing/500", nil)
}
