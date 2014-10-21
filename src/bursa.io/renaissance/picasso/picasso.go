package picasso

// The famous painter Picasso created many exceptional works. This package makes
// the same thing possible for rendering web server responses in Go.

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
)

// Allows us to not have to pass around the writer/request for every API call
type Picasso struct {
	writer  http.ResponseWriter
	Request *http.Request
}

func New(writer http.ResponseWriter, request *http.Request) *Picasso {
	p := &Picasso{Request: request}
	p.writer = writer
	return p
}

// The base function that you will call in most cases, handles loading, 404's,
// 500's and catching errors if they occur
func (self *Picasso) Render(view string, pass_ins map[string]string) {
	self.serveTemplate(view, pass_ins)
}

// Serves a 404 if the page cannot be found
func (self *Picasso) Serve404() {
	http.NotFound(self.writer, self.Request)
	return
}

// A 500 error if an unspecified error is thrown during the rendering process
func (self *Picasso) Serve500(err error) {
	log.Println(err.Error())
	http.Error(self.writer, http.StatusText(500), 500)
	return
}

// Handles Serving and Error Handling of Templates
func (self *Picasso) serveTemplate(view_name string, pass_ins map[string]string) {
	base := path.Join("static", "layout.html")
	view := path.Join("static", "views", view_name)

	// Handle 404's
	info, err := os.Stat(view)
	if err != nil {
		if os.IsNotExist(err) {
			self.Serve404()
		}
	}

	// If trying to render a directory
	if info.IsDir() {
		self.Serve404()
	}

	tmpl, err := template.ParseFiles(base, view)
	if err != nil {
		self.Serve500(err)
	}

	// If we made it this far, try rendering with passins
	err = tmpl.ExecuteTemplate(self.writer, "layout", pass_ins)
	if err != nil {
		self.Serve500(err)
	}
}
