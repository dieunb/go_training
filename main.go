package main

import (
	"log"
	"net/http"

	"github.com/dieunb/go_training/configs"
	"github.com/dieunb/go_training/pkg/controllers"
	"github.com/dieunb/go_training/pkg/middlewares"
	"github.com/gorilla/mux"
)

func main() {
	router := NewRouter()

	router.HandleFunc("/", middlewares.Logging(controllers.BookIndex)).Methods("GET")

	router.HandleFunc("/books/{title}/page/{page}", middlewares.Logging(controllers.BookShow)).Methods("GET")
	router.HandleFunc("/books/{title}", middlewares.Logging(controllers.BookDelete)).Methods("DELETE")

	err := http.ListenAndServe(":"+configs.Port(), router)

	if err != nil {
		log.Fatal("ListenAndServe Error: ", err)
	}
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Server CSS, JS & Images Statically.
	router.
		PathPrefix(configs.STATIC_DIR).
		Handler(http.StripPrefix(configs.STATIC_DIR, http.FileServer(http.Dir("."+configs.ASSETS_DIR))))
	return router
}
