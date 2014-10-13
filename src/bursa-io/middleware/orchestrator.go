package middleware

import (
	"net/http"
)

// An unofficial interface of sorts, allows us to create an array of Handler
// functions
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
