// This package handles the building of layouts, partials, and templates
// into a renderable package which can then be written to an HTTP response.
package picasso

import (
	"html/template"
	"net/http"
	"path"
	"path/filepath"

	"bursa.io/config"
	"runtime"
)

func Render(w http.ResponseWriter, layout string, view string, vars interface{}) {

	template_dir := getTemplateRoot()
	layout_path := path.Join(template_dir, layout+".tmpl")
	view_path := path.Join(template_dir, view+".tmpl")

	// partials := findPartials(layout_path)

	temp := template.Must(template.ParseFiles(layout_path, view_path))

	// Provides some visibility into template execution errors.
	if err := temp.Execute(w, vars); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Creates a relative path to our templates folder
func getTemplateRoot() string {
	_, filename, _, _ := runtime.Caller(1)
	filepath := path.Join(path.Dir(filename), "../../../")
	return path.Join(filepath, config.GetStringMapString("paths")["templates"])
}

// Searches the folder that the layout is defined in for a "/partials" folder
// and parses partial file names into a slice if the folder exists and it is
// populated
func findPartials(layout_path string) []string {

	expected_partial_dir := path.Join(path.Dir(layout_path), "partials")

	// now do a dir listing
	files, err := filepath.Glob(expected_partial_dir + "/*")
	if err != nil {
		files = make([]string, 2)
	}

	return files
}
