package picasso

import (
	"html/template"
	"log"
	"net/http"
	"path"

	"bursa.io/config"
)

func Render(w http.ResponseWriter, layout string, view string, vars interface{}) {

	// generate paths to our templates
	paths := config.GetStringMapString("paths")
	template_dir := paths["templates"]
	layout_path := path.Join(template_dir, layout+".tmpl")
	view_path := path.Join(template_dir, view+".tmpl")

	temp := template.Must(template.ParseFiles(layout_path, view_path))

	// Provides some visibility into template execution errors.
	if err := temp.Execute(w, vars); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
