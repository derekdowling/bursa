package controller

// This handles rendering our unauthenticated user facing static web pages.

import (
	"bursa.io/renaissance/picasso"
	"net/http"
)

type HomeController struct{}

// Returns the site home page
func (h *HomeController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	picasso := picasso.New(w, r)
	picasso.Render("index.html", nil)
}
