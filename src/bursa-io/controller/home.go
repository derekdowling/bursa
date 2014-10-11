package controller

import (
	"fmt"
	"html"
	"net/http"
)

type HomeController struct{}

func (h *HomeController) Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q \n", html.EscapeString(r.URL.Path))
}
