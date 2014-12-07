// This package handles all of our session setting logic within the app
package session

import (
	log "github.com/Sirupsen/logrus"
	"github.com/derekdowling/bursa/config"
	"github.com/gorilla/sessions"
	"net/http"
)

type Manager struct {
	Store *sessions.CookieStore
}

func New() *Manager {
	return &Manager{
		Store: getStore(),
	}
}

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
func (self *Manager) getAppSession(r *http.Request) *sessions.Session {

	// Retrieve or build a session from thr current request and return
	session, _ := self.Store.Get(r, getSessionName())
	return session
}

// Handles creating a new user session
func (self *Manager) CreateUserSession(w http.ResponseWriter, r *http.Request, user_id int64) {

	session := self.getAppSession(r)
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   300,
		HttpOnly: true,
	}

	session.Values["user_id"] = user_id
	err := session.Save(r, w)
	if err != nil {
		log.Fatal(err)
	}
}

// Checks whether or not a user is already logged in via their session token
func (self *Manager) GetUserId(r *http.Request) int64 {
	session := self.getAppSession(r)

	user_id := session.Values["user_id"]
	return user_id.(int64)
}
