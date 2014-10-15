package picasso

import (
	"bursa-io/logger"
	"html/template"
	"net/http"
	"os"
	"path"
)

// Allows us to not have to pass around the writer/request for every API call
type Picasso struct {
	writer  http.ResponseWriter
	Request *http.Request
}

func NewPicasso(w http.ResponseWriter, r *http.Request) *Picasso {
	p := &Picasso{Request: r}
	p.writer = w
	return p
}

// The base function that you will call in most cases, handles loading, 404's,
// 500's and catching errors if they occur
func (p *Picasso) Render(view string, pass_ins map[string]string) {
	p.serveTemplate(view, pass_ins)
}

// Serves a 404 if the page cannot be found
func (p *Picasso) Serve404() {
	http.NotFound(p.writer, p.Request)
	return
}

// A 500 error if an unspecified error is thrown during the rendering process
func (p *Picasso) Serve500(err error) {
	logger.Println(err.Error())
	http.Error(p.writer, http.StatusText(500), 500)
	return
}

// Handles loading, combining, and building template views
func (p *Picasso) serveTemplate(view_name string, pass_ins map[string]string) {
	base := path.Join("static", "layout.html")
	view := path.Join("static", "views", view_name)

	// Handle 404's
	info, err := os.Stat(view)
	if err != nil {
		if os.IsNotExist(err) {
			p.Serve404()
		}
	}

	// If trying to render a directory
	if info.IsDir() {
		p.Serve404()
	}

	tmpl, err := template.ParseFiles(base, view)
	if err != nil {
		p.Serve500(err)
	}

	// If we made it this far, try rendering with passins
	err = tmpl.ExecuteTemplate(p.writer, "layout", pass_ins)
	if err != nil {
		p.Serve500(err)
	}
}
