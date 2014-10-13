package controller

import (
	"bursa-io/middleware"
	"fmt"
	"html"
	"net/http"
)

type WalletController struct{}

func (wc *WalletController) GetHandler() middleware.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q \n", html.EscapeString(r.URL.Path))
	}
}
