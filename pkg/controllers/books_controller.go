package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/dieunb/go_training/pkg/models"
	"github.com/gorilla/mux"
)

func BookIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("layout.html"))

	data := models.BookIndex{
		PageTitle: "My books list",
		Books: []models.Book{
			{Title: "Ruby"},
			{Title: "Ruby on Rails"},
			{Title: "Go"},
		},
	}
	err := tmpl.Execute(w, data)

	if err != nil {
		log.Fatal("Render error ", err)
	}
}

func BookShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]

	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
}

func BookDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	fmt.Fprintf(w, "Book %s have deleted!", title)
}
