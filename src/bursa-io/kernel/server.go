package main

import (
	"bursa-io/middleware"
	"fmt"
	"github.com/gorilla/mux"
	"html"
	"net/http"
)

func main() {
	route()
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q \n", html.EscapeString(r.URL.Path))
}

func walletHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Wallets are here!, %q \n", html.EscapeString(r.URL.Path))
}

// Handles our basic routes
// http://www.gorillatoolkit.org/pkg/mux
func route() {
	router := mux.NewRouter()

	router.HandleFunc("/", middleware.GlobalHandler.WithController(homeHandler))

	// Just some basic other examples
	router.HandleFunc("/wallets/create", middleware.GlobalHandler.WithController(walletHandler))
	router.HandleFunc("/wallets/{id:[0-9]+", homeHandler).Methods("GET")

	// Pass our router to net/http
	http.Handle("/", router)

	// Listen and serve on localhost:8080
	http.ListenAndServe(":8080", nil)
}
