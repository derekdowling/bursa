package picasso

import (
	"bursa.io/config"
	"html/template"
	"net/http"
	"path"
)

func Render(w http.ResponseWriter, layout string, view string, vars interface{}) {

	// generate paths to our templates
	paths := config.GetStringMapString("paths")
	template_dir := paths["templates"]
	layout_path := path.Join(template_dir, layout+".tmpl")
	view_path := path.Join(template_dir, view+".tmpl")

	temp := template.Must(template.New("view").ParseFiles(layout_path, view_path))

	temp.ExecuteTemplate(w, layout, vars)
}
