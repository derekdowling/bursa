package main

// This is essentially our server kernel. It handles
import (
	"bursa-io/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	route()
}

// Handles our basic routes
// http://www.gorillatoolkit.org/pkg/mux
func route() {
	router := mux.NewRouter()

	router.HandleFunc("/", Orchestartor.WithController(homeController))

	// Just some basic other examples
	router.HandleFunc("/wallets/create", Orchestrator.WithController(walletController))
	router.HandleFunc("/wallets/{id:[0-9]+", Orchestrator.WithController(walletController)).Methods("GET")

	// Pass our router to net/http
	http.Handle("/", router)

	// Listen and serve on localhost:8080
	http.ListenAndServe(":8080", nil)
}

func init() {
	Orchestrator := new(ControllerController)

	config := new(ConfigMiddleware)
	Orchestrator.AddMiddleware(config.GetHandler())
}
