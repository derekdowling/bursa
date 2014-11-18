package main

import (
	_ "github.com/derekdowling/bursa/kernel"
	log "github.com/Sirupsen/logrus"
	"github.com/derekdowling/bursa/config"
	"github.com/derekdowling/bursa/kernel"
	"github.com/derekdowling/bursa/controller/home"
	"github.com/derekdowling/bursa/controller/app"
	"github.com/codegangsta/negroni"
	"flag"
	"github.com/gorilla/mux"
	"net/http"
)

var production = flag.Bool("prod", false, "Starts the server in production mode")

func main() {
	// Parse in flags
	flag.Parse()

	// starts the kernel in production mode
	Start(*production)
}

// This handles starting up our web kernel. It'll load our routes, controllers, and
// middleware.
func Start(production bool) {

	// get our stack rolling
	stack := kernel.BuildStack(production)

	registerRouter(stack)

	// figure out what port we need to be on
	port := config.Server.GetStringMapString("ports")["http"]

	// output to help notify that the server is loaded
	log.WithFields(log.Fields{"port": port}).Info("Ready for requests with:")

	// start and log server output
	log.Fatal(http.ListenAndServe(port, stack))

	// TODO: get below working when we want HTTPS in prod
	// Listen, Serve, Log
	// log.Fatal(
	// http.ListenAndServeTLS(
	// config.Server.GetString("server.Https_Port"),
	// "src/bursa.io/server/certs/cert.pem",
	// "src/bursa.io/server/certs/key.pem",
	// stack,
	// ),
	// )
}

func registerRouter(stack *negroni.Negroni) {

	// Builds our router and gives it routes
	router := buildRouter()

	// Serve static assets that the website requests
	static_routes := config.Server.GetStringMapString("static_routes")

	for url, local := range static_routes {

		log.WithFields(log.Fields{
			"route": url,
			"path":  local,
		}).Info("Asset Path:")

		router.PathPrefix(url).Handler(
			http.FileServer(http.Dir(local)),
		)
	}

	stack.UseHandler(router)
}

// Builds our routes
// http://www.gorillatoolkit.org/pkg/mux
func buildRouter() *mux.Router {

	// Create a Gorilla Mux Router
	router := mux.NewRouter()

	router.Queries("email", "")

	// Website Routes
	router.HandleFunc("/", home.HandleIndex).Methods("GET")
	router.HandleFunc("/app", app.HandleIndex).Methods("GET")
	router.HandleFunc("/about", home.HandleAbout).Methods("GET")
	router.HandleFunc("/login", home.HandleLogin).Methods("GET")
	router.HandleFunc("/login", home.HandlePostLogin).Methods("POST")
	router.HandleFunc("/signup", home.HandleSignup).Methods("GET")
	router.HandleFunc("/signup", home.HandlePostSignup).Methods("POST")
	router.HandleFunc("/forgot-password", home.HandleForgotPassword).Methods("GET")
	// router.HandleFunc("/forgot-password", home.HandlePostSignup).Methods("POST").Queries("email", "")

	// Our 404 Handler
	router.NotFoundHandler = http.HandlerFunc(home.Handle404)

	// API Routes
	// TODO: Rest Layer

	return router
}
