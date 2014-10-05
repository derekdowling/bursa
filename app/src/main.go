package bursa

import (
  "net/http"
  "fmt"
  "html"
  "github.com/gorilla/mux"
  "github.com/derekdowling/bursa/models"
)

func main() {
    route()
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

// Handles our basic routes
// http://www.gorillatoolkit.org/pkg/mux
func route() {
	router := mux.NewRouter()

	router.HandleFunc("/", homeHandler)

	// Just some basic other examples
	router.HandleFunc("/wallets/create", homeHandler)
  router.HandleFunc("/wallets/{id:[0-9]+", homeHandler).Methods("GET")

  // Pass our router to net/http
  http.Handle("/", router)

  // Listen and serve on localhost:8080
  http.ListenAndServe(":8080", nil)
}
