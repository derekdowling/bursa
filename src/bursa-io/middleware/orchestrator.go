package middleware

import (
	"bursa-io/picasso"
	"net/http"
)

// An unofficial interface of sorts, allows us to create an array of Handler
// functions
type MuxHandler func(w http.ResponseWriter, r *http.Request)
type Handler func(p *picasso.Picasso)

type ControllerController struct {
	handlers []Handler
}

func (self *ControllerController) AddMiddleware(handler Handler) {
	self.handlers = append(self.handlers, handler)
}

// Questions:
// Public methods - do they need uppercase?
func (self *ControllerController) WithController(controller Handler) MuxHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		pablo := picasso.NewPicasso(w, r)
		for _, handler := range self.handlers {
			handler(pablo)
		}
		controller(pablo)
	}
}
