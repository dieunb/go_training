package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

const (
	ASSETS_DIR = "/assets/"
	STATIC_DIR = "/static/"
	PORT       = "4000"
)

type Book struct {
	Title string
}

type BookToRead struct {
	PageTitle string
	Books     []Book
}

func main() {
	router := NewRouter()

	router.HandleFunc("/", BookToReadHandler).Methods("GET")

	router.HandleFunc("/books/{title}/page/{page}", BookHandler).Methods("GET")
	router.HandleFunc("/books/{title}", DeleteBookHandler).Methods("DELETE")

	err := http.ListenAndServe(":"+PORT, router)

	if err != nil {
		log.Fatal("ListenAndServe Error: ", err)
	}
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Server CSS, JS & Images Statically.
	router.
		PathPrefix(STATIC_DIR).
		Handler(http.StripPrefix(STATIC_DIR, http.FileServer(http.Dir("."+ASSETS_DIR))))
	return router
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
	err := tmpl.Execute(w, data)

	if err != nil {
		log.Fatal("Render error ", err)
	}
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
