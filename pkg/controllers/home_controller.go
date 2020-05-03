package controllers

import (
	"net/http"

	"github.com/dieunb/go_training/pkg/views"
)

func HomeIndex(w http.ResponseWriter, r *http.Request) {
	data := struct {
		PageTitle string
		Name      string
	}{
		"Hello World!",
		"Justin",
	}

	views.Render("home/index.html", data, &w)
}
