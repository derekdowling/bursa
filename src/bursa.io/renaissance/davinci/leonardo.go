package davinci

// This middleware manager is name in tribute of the Great Leonardo da Vinci. The
// goal of this package is designed to allow you to create your own truly remarkable
// contraptions without being limited by how you are going to put them all together.

import (
	"bursa-io/picasso"
	"net/http"
)

// A blueprint allows Davinci to interface with standard Go http middleware
type Blueprint func(w http.ResponseWriter, r *http.Request)

// A mechanism is one of the middleware pieces in our invention. We pass Picasso
// into
type Mechanism func(p *picasso.Picasso)

// All of the mechanisms put together makes a contraption
// Can optionally specify a normalizer that allows you to pass in additional objects
// into your mechanisms/controllers
type Contraption struct {
	mechanisms []Mechanism
	normalizer Blueprint
}

func (self *Contraption) SetNormalizer(normalizer Blueprint) {

}

// Adds a new mechanism(Middleware) to this current contraption
func (self *Contraption) AddMechanism(m Mechanism) {
	self.mechanisms = append(self.mechanisms, m)
}

// Allows 3rd Party Blueprints to be used without ruining the contraption
func (self *Contraption) ImportMechanism(blueprint Blueprint, p *picasso.Picasso) {
	mechanism := func(p *picasso.Picasso) {
		blueprint(p.Critique())
	}
	self.AddMechanism(mechanism)
}

// Puts the contraption together and provides a controller to make it useful
func (self *Contraption) Create(controller Mechanism) Blueprint {
	return func(w http.ResponseWriter, r *http.Request) {
		pablo := picasso.NewPicasso(w, r)
		for _, mechanism := range self.mechanisms {
			mechanism(pablo)
		}
		controller(pablo)
	}
}
