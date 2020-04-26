package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

type Book struct {
	Title string
}

type BookToRead struct {
	PageTitle string
	Books     []Book
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", BookToReadHandler).Methods("GET")

	r.HandleFunc("/books/{title}/page/{page}", BookHandler).Methods("GET")
	r.HandleFunc("/books/{title}", DeleteBookHandler).Methods("DELETE")

	http.ListenAndServe(":4000", r)
}

func BookToReadHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("layout.html"))

	data := BookToRead{
		PageTitle: "My books list",
		Books: []Book{
			{Title: "Ruby"},
			{Title: "Ruby on Rails"},
			{Title: "Go"},
		},
	}
	tmpl.Execute(w, data)
}

func BookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]

	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
}

func DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	fmt.Fprintf(w, "Book %s have deleted!", title)
}
