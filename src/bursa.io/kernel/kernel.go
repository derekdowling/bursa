package kernel

// The Kernel is what builds and runs our webserver. Here our middleware, routes,
// and controllers are all defined and assembled. Our configuration files are also
// loaded into Viper so they can be used from anywhere after the Kernel loads.

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

// This handles starting up our web kernel. It'll load our routes, controllers, and
// middleware.
func start() {

	// Builds our router and gives it routes
	router := buildRouter()

	// Serve static assets that the website requests
	fs := http.FileServer(http.Dir("static"))
	router.Handle("/static/", http.StripPrefix("/static/", fs))

	// Pass our router to net/http
	http.Handle("/", router)

	// Listen and serve on localhost:8080
	http.ListenAndServe(":8080", nil)
}

// Builds our routes
// http://www.gorillatoolkit.org/pkg/mux
func buildRouter() *mux.Router {

	// Create a Gorilla Mux Router
	router := mux.NewRouter()

	// Get our mapped routes
	routes := defineRoutes()

	// Add them to the router
	for route, handler := range routes {
		router.Handle(route, handler)
	}

	return router
}

// Initializes routes for the router
func defineRoutes() map[string]http.Handler {

	routes := make(map[string]http.Handler)

	//prepare middleware
	middleware := defineMiddleware()

	// Initialize Controllers Here
	walletController := new(controller.WalletController)
	homeController := new(controller.HomeController)

	// Website Routes
	routes["/"] = CreateBlueprint.AddMechanisms(mechanisms).AddController(home)
	routes["/wallet/create"] = CreateBlueprint.AddMechanisms(mechanisms).AddController(controller.WalletController)

	return routes
}

// Defines the mechanisms that we will be using and returns them
func defineMiddleware() []Mechanism {
	middleware := []Mechanism{}
	return middleware
}
