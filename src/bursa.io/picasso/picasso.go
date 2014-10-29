package picasso

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"os"

	"bursa.io/config"
)

func Render(w http.ResponseWriter, layout string, view string, vars interface{}) {

	// generate paths to our templates
	paths := config.GetStringMapString("paths")
	template_dir := paths["templates"]
	layout_path := path.Join(template_dir, layout+".tmpl")
	view_path := path.Join(template_dir, view+".tmpl")

	log.Println(view_path)
	log.Println(layout_path)
	if _, err := os.Stat(layout_path); err != nil {
		log.Fatalf("Damnit", err)
	}
	if _, err := os.Stat(view_path); err != nil {
		log.Fatalf("Damnit", err)
	}

	temp := template.Must(template.New("view").ParseFiles(layout_path))

	log.Println(temp)

	temp.ExecuteTemplate(w, layout, vars)
}
