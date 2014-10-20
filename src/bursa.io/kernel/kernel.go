package kernel

// The Kernel is what builds and runs our webserver. Here our middleware, routes,
// and controllers are all defined and assembled. Our configuration files are also
// loaded into Viper so they can be used from anywhere after the Kernel loads.

import (
	"bursa.io/config"
	"bursa.io/controller"
	"bursa.io/renaissance/davinci"
	"github.com/gorilla/mux"
	"github.com/unrolled/secure"
	"net/http"
)

func init() {
	// loads our config into Viper so it can be used anywhere
	config.LoadConfig()
}

// This handles starting up our web kernel. It'll load our routes, controllers, and
// middleware.
func start(production bool) {

	// Builds our router and gives it routes
	router := buildRouter()

	if production {
		// Serve static assets that the website requests
		fs := http.FileServer(http.Dir("static"))
		router.Handle("/static/", http.StripPrefix("/static/", fs))
	}

	// Build our contraption middleware and add the router
	// as the last piece
	app := new(Contraption)
	app.AddSet(buildMiddleware())
	app.Add(router)

	// Listen, Serve, Log
	log.Fatal(http.ListenAndServeTLS(viper.GetString("server.Https_Port"), "cert.pem", "key.pem", app))
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

// Sets our secure middleware based on what mode we are in
func createSecureMiddleware(production bool) {
	secureMiddleware := secure.New(secure.Options{
		AllowedHosts:          viper.GetStringSlice("Allowed_Hosts"),
		SSLRedirect:           true,
		SSLHost:               viper.GetString("server.SSL_Host"),
		SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"},
		STSSeconds:            315360000,
		STSIncludeSubdomains:  true,
		FrameDeny:             true,
		ContentTypeNosniff:    true,
		BrowserXssFilter:      true,
		ContentSecurityPolicy: "default-src 'self'",
	})
}
