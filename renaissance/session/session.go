// This package handles all of our session setting logic within the app
package session

import (
	"github.com/derekdowling/bursa/config"
	"github.com/gorilla/sessions"
	"net/http"
)

// Interface for getting a session store. Overwrite this to Mock in testing.
func getStore() *sessions.CookieStore {
	return sessions.NewCookieStore([]byte(getSessionKey()))
}

func getSessionKey() string {
	return config.App.GetString("session.key")
}

func getSessionName() string {
	return config.App.GetString("session.name")
}

// Returns the current sessions assoicated with the request, returns a blank
// session if one isn't setup
func getAppSession(r *http.Request) *sessions.Session {
	store := getStore()
	session, _ := store.Get(r, getSessionName())
	return session
}

// Handles creating a new user session
func CreateUserSession(w http.ResponseWriter, r *http.Request, user_id int64) {

	session := getAppSession(r)
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   300,
		HttpOnly: true,
	}

	session.Values["user_id"] = user_id
	session.Save(r, w)
}

// Checks whether or not a user is already logged in via their session token
func GetUserId(r *http.Request) int64 {
	session := getAppSession(r)

	user_id := session.Values["user_id"]
	return user_id.(int64)
}
