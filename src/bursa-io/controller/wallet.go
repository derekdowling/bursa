package controller

import (
	"fmt"
	"html"
	"net/http"
)

type WalletController struct{}

func (wc *WalletController) Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q \n", html.EscapeString(r.URL.Path))
}
