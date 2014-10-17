package controller

// This handles rendering our unauthenticated user facing static web pages.

import (
	"bursa.io/renaissance/satchel"
	"net/http"
)

type HomeController struct{}

// Returns the site home page
func (h *HomeController) Trigger(s *satchel.Satchel) middleware.Handler {
	return func(p *Picasso) {
		picasso.Render("index.html", nil)
	}
}
