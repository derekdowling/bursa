package davinci

// This middleware manager is name in tribute of the Great Leonardo da Vinci. The
// goal of this package is to enable you to easily apply middleware layers
// to your web contraptions
import (
	"bursa.io/renaissance/satchel"
	"log"
	"net/http"
)

// A little struct that lets us easily
type HTTPHandler struct {
	handler http.Handler
}

// Allows us to trigger our wrapper for html.Handler compatible
// middleware
func (self *HTTPHandler) Trigger(s *satchel.Satchel) {
	self.handler.ServeHTTP(s.Context())
}

// Call this on a third-party middelware handler function to return a normalized
// Mechanism that DaVinci can use
func normalizeHandler(handler http.Handler) *HTTPHandler {
	mechanism := new(HTTPHandler)
	mechanism.handler = handler
	return mechanism
}

// Acts as a really basic interface for our DaVinci components so we can invoke
// them in a sane manner.
type Mechanism interface {
	Trigger(s *satchel.Satchel)
}

// Stores all of the middleware Mechanisms each handler should have
type Blueprint struct {
	mechanisms []Mechanism
	controller Mechanism
}

// Contructor for a fluent interface
func StartBlueprint() *Blueprint {
	return new(Blueprint)
}

// Adds a new mechanism(Middleware) to this current contraption
// if the mechanism is a Golang default, we convert it automatically
// Might be a little slow since it is a run time mechanism, but until we
// hit scale we'll have bigger fish to fry
func (self *Blueprint) AddMechanisms(mechanisms []Mechanism) *Blueprint {
	self.mechanisms = mechanisms
	return self
}

// Adds a controller and then builds the Contraption from the Blueprint. Is careful
// to create a new Contraption with unique middleware layers so we don't run over
// pointer values
func (self *Blueprint) AddController(controller Mechanism) *Blueprint {

	if self.controller != nil {
		log.Panic("This Blueprint already has a controller, cannot define another")
	}

	self.controller = controller
	return self
}

// Builds a new Contraption from a blueprint, taking care to create new instances
// of all structs
func (self *Blueprint) build() *Contraption {
	// create a new Contraption with pre-initialized data
	contraption := new(Contraption)

	// iterate through middleware Mechanisms and create new instances of each so
	// each contraption is isolated
	for _, mechanism := range self.mechanisms {
		clone := mechanism
		contraption.mechanisms = append(contraption.mechanisms, clone)
	}

	// Now add the Controller to the list of Mechanisms
	clone := self.controller
	contraption.mechanisms = append(contraption.mechanisms, clone)

	return contraption
}

// Implements the http.Handler function so that you can pass your Mux
// a Contraption pointer directly
func (self *Blueprint) serveHTTP(w http.ResponseWriter, r *http.Request) {

	// creates a new Contraption from a Blueprint and run it
	contraption := self.build()
	contraption.Run(satchel.New(w, r))
}

type Contraption struct {
	mechanisms []Mechanism
}

// Runs the contraption by triggering it's Mechanisms in sequential order
func (self *Contraption) Run(s *satchel.Satchel) {
	for _, mechanism := range self.mechanisms {
		mechanism.Trigger(s)
	}
}
