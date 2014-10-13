package controller

import (
	"bursa-io/middleware"
	"fmt"
	"html"
	"net/http"
)

type HomeController struct{}

func (h *HomeController) GetHandler() middleware.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome home, %q \n", html.EscapeString(r.URL.Path))
	}
}
