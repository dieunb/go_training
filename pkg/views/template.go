package views

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

var (
	dir, _ = os.Getwd()
	header = filepath.Join(dir, "web", "template", "shared", "header.html")
	footer = filepath.Join(dir, "web", "template", "shared", "footer.html")
	layout = filepath.Join(dir, "web", "template", "layouts", "application.html")
)

func Render(file_name string, data interface{}, w *http.ResponseWriter) {
	view_file_path := filepath.Join(dir, "web", "template", file_name)
	tmpl, err := template.ParseFiles(layout, header, footer, view_file_path)

	if err != nil {
		http.Error(*w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(*w, data); err != nil {
		http.Error(*w, err.Error(), http.StatusInternalServerError)
		return
	}
}
