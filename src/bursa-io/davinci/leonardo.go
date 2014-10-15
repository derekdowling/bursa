package davinci

// This middleware manager is name in tribute of the Great Leonardo da Vinci. The
// goal of this package is designed to allow you to create your own truly remarkable
// contraptions without being limited by how you are going to put them all together.

import (
	"bursa-io/picasso"
	"net/http"
)

// A blueprint allows us to interface with the inventions of others after we
// "Discover" them ourselves.
type Blueprint func(w http.ResponseWriter, r *http.Request)

// A mechanism is one of the middleware pieces on our invention.
type Mechanism func(p *picasso.Picasso)

// All of the mechanisms put together in the order they were created
type Contraption struct {
	mechanisms []Mechanism
}

// Adds a new mechanism to the overall contraption
func (self *Contraption) AddMechanism(m Mechanism) {
	self.mechanisms = append(self.mechanisms, m)
}

// Allows DaVinci to use 3rd Party Middleware without ruining one of his
// contraptions.
func (self *Contraption) AddDiscovery(blueprint Blueprint, p *picasso.Picasso) {
	mechanism := func(p *picasso.Picasso) {
		blueprint(p.Critique())
	}
	self.AddMechanism(mechanism)
}

// Adds the final touch, a control to the contraption so that it becomes useful
// and then returns a blueprint so meer mortals can build and operate it
func (self *Contraption) Invent(controller Mechanism) Blueprint {
	return func(w http.ResponseWriter, r *http.Request) {
		pablo := picasso.NewPicasso(w, r)
		for _, mechanism := range self.mechanisms {
			mechanism(pablo)
		}
		controller(pablo)
	}
}
