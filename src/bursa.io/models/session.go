package models

// Handles all session interactions

import (
	"github.com/gorilla/sessions"
)

func loadStore() sessions.Store {
	// Load our session store
	store := sessions.NewCookieStore([]byte(config.GetString("session-key")))
	return store
}

// Handles creating a new user session
func CreateUserSession() {

	session, _ := loadStore(r, "app-session")

	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   300,
		HttpOnly: true,
	}

	session.Values[LoggedIn] = true
	session.Save(r, w)
}

// Checks whether or not a user is already logged in via their session token
func LoggedIn(r *http.Request) bool {
	session, _ := store.Get(r, "app-session")

	logged_in = session.Get(LoggedIn)
	return logged_in
}
