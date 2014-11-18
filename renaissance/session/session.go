package session

// Handles all session interactions

import (
	"github.com/derekdowling/bursa/config"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
)

func loadStore() sessions.Store {
	// Load our session store
	store := sessions.NewCookieStore([]byte(config.Server.GetString("session.Key.Main")))
	return store
}

func getAppSession(r *http.Request) *sessions.Session {
	store := loadStore()

	// always returns a blank session if none is present
	session, _ := store.Get(r, config.Server.GetString("session.Key.App"))
	return session
}

// Handles creating a new user session
func CreateUserSession(w http.ResponseWriter, r *http.Request) {

	session := getAppSession(r)

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
	session := getAppSession(r)

	log.Print(session)
	logged_in := session.Values[LoggedIn]
	log.Print(logged_in)
	return logged_in.(bool)
}
