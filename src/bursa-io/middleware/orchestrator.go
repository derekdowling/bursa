package middleware

import (
	"net/http"
)

// An unofficial interface of sorts, allows us to create an array of Handler
// functions
type Handler func(w http.ResponseWriter, r *http.Request)

var Orchestrator = new(ControllerController)

type ControllerController struct {
	handlers []Handler
}

func (self *ControllerController) AddMiddleware(handler Handler) {
	self.handlers = append(self.handlers, handler)
}

// Questions:
// Public methods - do they need uppercase?
func (self *ControllerController) WithController(controller Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, handler := range self.handlers {
			handler(w, r)
		}
		controller(w, r)
	}
}

// This is where we add all of our middleware layers in

// Gets called when the package is first loaded, add middleware here
func init() {

	config := new(ConfigMiddleware)

	Orchestrator.AddMiddleware(config.GetHandler())
}
