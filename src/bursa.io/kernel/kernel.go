package kernel

// The Kernel is what builds and runs our webserver. Here our middleware, routes,
// and controllers are all defined and assembled. Our configuration files are also
// loaded into Viper so they can be used from anywhere after the Kernel loads.

import (
	"bursa.io/config"
	"bursa.io/controller/app"
	"bursa.io/controller/home"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/unrolled/secure"
	"log"
	"net/http"
	"os"
)

func init() {

	// set log output path
	log.SetOutput(os.Stdout)

	// loads our config into Viper so it can be used anywhere
	config.LoadConfig()
}

// This handles starting up our web kernel. It'll load our routes, controllers, and
// middleware.
func Start(production bool) {

	// Build our contraption middleware and add the router
	// as the last piece
	stack := negroni.New()

	// define our list of production middleware here for now
	if production {
		// Secure middleware has a Negroni integration, hence the wonky syntax
		stack.Use(negroni.HandlerFunc(secureMiddleware().HandlerFuncWithNext))
	} else {
		stack.Use(negroni.NewLogger())
	}

	// Builds our router and gives it routes
	router := buildRouter()

	// Serve static assets that the website requests
	static_routes := config.GetStringMapString("static_routes")
	log.Println("Loading static assets:")
	for url, local := range static_routes {
		log.Println("route:", url, "- path:", local)
		router.PathPrefix(url).Handler(
			http.FileServer(http.Dir(local)),
		)
	}

	stack.UseHandler(router)

	// port := config.GetString("server.Ports.Http")
	port := viper.GetStringMapString("ports")["http"]
	log.Printf("Listening for requests on %s", port)
	log.Fatal(http.ListenAndServe(port, stack))
	// Listen, Serve, Log
	// log.Fatal(
	// http.ListenAndServeTLS(
	// config.GetString("server.Https_Port"),
	// "src/bursa.io/server/certs/cert.pem",
	// "src/bursa.io/server/certs/key.pem",
	// stack,
	// ),
	// )
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
		router.HandleFunc(route, handler)
	}

	return router
}

// Initializes routes for the router
func defineRoutes() map[string]http.HandlerFunc {

	routes := make(map[string]http.HandlerFunc)

	// Website Routes
	routes["/"] = home.HandleIndex
	routes["/app"] = app.HandleIndex

	return routes
}

// Sets our secure middleware based on what mode we are in
func secureMiddleware() *secure.Secure {
	secureMiddleware := secure.New(secure.Options{
		AllowedHosts:          config.GetStringSlice("server.Allowed_Hosts"),
		SSLRedirect:           true,
		SSLHost:               config.GetString("server.SSL_Host"),
		SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"},
		STSSeconds:            315360000,
		STSIncludeSubdomains:  true,
		FrameDeny:             true,
		ContentTypeNosniff:    true,
		BrowserXssFilter:      true,
		ContentSecurityPolicy: "default-src 'self'",
	})
	return secureMiddleware
}
