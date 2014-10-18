package satchel

// The trusty renaissance satchel holds onto all those things you value most and
// allows you to access them intelligent ways

import (
	"bursa.io/renaissance/picasso"
	"net/http"
)

type Satchel struct {
	writer  http.ResponseWriter
	Request *http.Request
}

func New(w http.ResponseWriter, r *http.Request) *Satchel {
	satchel := new(Satchel)
	satchel.writer = w
	satchel.Request = r
	return satchel
}

func (self *Satchel) GetPicasso() *picasso.Picasso {
	return picasso.NewPicasso(self.writer, self.Request)
}

func (self *Satchel) Context() (http.ResponseWriter, *http.Request) {
	return self.writer, self.Request
}
