package main

// This is essentially our server kernel. It handles
import (
	"bursa-io/config"
	"bursa-io/controller"
	"bursa-io/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

func init() {
	// loads our config into Viper so it can be used anywhere
	config.LoadConfig()
}

func main() {
	route()
}

// Handles our basic routes
// http://www.gorillatoolkit.org/pkg/mux
func route() {

	// Define and populate our middleware layers
	contraption := new(Contraption)

	// Initialize Controllers Here
	walletController := new(controller.WalletController)
	homeController := new(controller.HomeController)

	// Setup Routes
	router := mux.NewRouter()
	router.HandleFunc("/", contraption.Invent(homeController.GetHandler()))
	router.HandleFunc("/wallets/create", contraption.Invent(walletController.GetHandler()))
	router.HandleFunc("/wallets/{id:[0-9]+", contraption.Invent(walletController.GetHandler())).Methods("GET")

	// Serve static assets that the website requests
	fs := http.FileServer(http.Dir("static"))
	router.Handle("/static/", http.StripPrefix("/static/", fs))

	// Pass our router to net/http
	http.Handle("/", router)

	// Listen and serve on localhost:8080
	http.ListenAndServe(":8080", nil)
}
