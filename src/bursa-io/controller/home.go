package controller

// This handles rendering our unauthenticated user facing static web pages.

import (
	"bursa-io/middleware"
	"bursa-io/picasso"
	"net/http"
)

type HomeController struct{}

// Returns the site home page
func (h *HomeController) GetHandler() middleware.Handler {
	return func(p *Picasso) {
		picasso.Render("index.html", nil)
	}
}
