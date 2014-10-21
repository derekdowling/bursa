// This firewall handes basic authorization checking. If the user is not authorized
// it redirects them back to safety where they can auth
package firewall

import (
	"bursa.io/config"
	"bursa.io/renaissance/session"
	"net/http"
)

// Adds our middleware interface
func ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	authorized := session.LoggedIn(r)

	// If not authorized, redirect to login
	if !authorized {
		http.Redirect(w, r, config.GetString("app.Login_Url"), http.StatusUnauthorized)
	}

	// if the user is authorized proceed
	next(w, r)
}
