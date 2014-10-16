package main

// This is essentially our server kernel. It handles
import (
	"bursa.io/config"
	"bursa.io/controller"
	"bursa.io/renaissance/davinci"
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
	mechanisms := []Mechanisms{}

	// Initialize Controllers Here
	walletController := new(controller.WalletController)
	homeController := new(controller.HomeController)

	// Setup Routes
	router := mux.NewRouter()
	router.HandleFunc("/", CreateBlueprint.AddMechanisms(mechanisms).AddController(home))
	router.HandleFunc("/wallets/create", CreateBlueprint.AddMechanisms(mechanisms).AddController(wallet))
	// Serve static assets that the website requests
	fs := http.FileServer(http.Dir("static"))
	router.Handle("/static/", http.StripPrefix("/static/", fs))

	// Pass our router to net/http
	http.Handle("/", router)

	// Listen and serve on localhost:8080
	http.ListenAndServe(":8080", nil)
}
