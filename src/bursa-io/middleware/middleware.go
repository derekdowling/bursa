package middleware

import (
	"fmt"
	"net/http"
)

type Handler func(w http.ResponseWriter, r *http.Request)

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

var GlobalHandler = new(ControllerController)

func init() {
	GlobalHandler.AddMiddleware(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "DERP \n")
	})

	GlobalHandler.AddMiddleware(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "WOOT \n")
	})
}
