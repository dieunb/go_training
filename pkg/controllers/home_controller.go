package controllers

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func HomeIndex(w http.ResponseWriter, r *http.Request) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	header := filepath.Join(dir, "web", "template", "shared", "header.html")
	footer := filepath.Join(dir, "web", "template", "shared", "footer.html")
	layout := filepath.Join(dir, "web", "template", "layouts", "application.html")
	home := filepath.Join(dir, "web", "template", "home", "index.html")

	tmpl, err := template.ParseFiles(layout, header, footer, home)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		PageTitle string
		Name      string
	}{
		"Hello World!",
		"Justin",
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
