package kernel

// The Kernel is what builds and runs our webserver. Here our middleware, routes,
// and controllers are all defined and assembled. Our configuration files are also
// loaded into Viper so they can be used from anywhere after the Kernel loads.

import (
	log "github.com/Sirupsen/logrus"
	"github.com/derekdowling/bursa/config"
	"github.com/derekdowling/bursa/middleware/logger"
	"github.com/derekdowling/bursa/middleware/logtext"
	"github.com/unrolled/secure"
	"github.com/codegangsta/negroni"
	"os"
)

func init() {

	// loads our config into Viper so it can be used anywhere
	config.LoadConfig()

	log_mode := config.Server.GetStringMapString("logging")["mode"]
	if log_mode == "production" {
		// Log as JSON instead of the default ASCII formatter.
		log.SetFormatter(&log.JSONFormatter{})
		// log.SetOutput(logstash)
	} else {

		log.SetLevel(log.DebugLevel)

		// gives our logger file/line/stack traces
		log.SetFormatter(logtext.NewLogtext(new(log.TextFormatter), false))
	}

	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(os.Stderr)
}


// Handle's putting the whole stack together
func BuildStack(production bool) *negroni.Negroni {
	// Build our contraption middleware and add the router
	// as the last piece
	stack := negroni.New()

	// define our list of production middleware here for now
	if production {
		// Turns on production API Keys
		config.Server.Set("production", true)
		// Secure middleware has a Negroni integration, hence the wonky syntax
		stack.Use(negroni.HandlerFunc(secureMiddleware().HandlerFuncWithNext))
	} else {
		stack.Use(logger.NewLogger())
	}

	return stack
}

// Sets our secure middleware based on what mode we are in
func secureMiddleware() *secure.Secure {
	secureMiddleware := secure.New(secure.Options{
		AllowedHosts:          config.Server.GetStringSlice("server.Allowed_Hosts"),
		SSLRedirect:           true,
		SSLHost:               config.Server.GetString("server.SSL_Host"),
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
