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

	router.HandleFunc("/", middlewares.Chain(controllers.HomeIndex, middlewares.Method("GET"), middlewares.Logging())).Methods("GET")

	router.HandleFunc("/books/{title}/page/{page}", middlewares.Chain(controllers.BookShow, middlewares.Method("GET"), middlewares.Logging())).Methods("GET")

	router.HandleFunc("/books/{title}", middlewares.Chain(controllers.BookDelete, middlewares.Method("DELETE"), middlewares.Logging())).Methods("DELETE")

	router.HandleFunc("/api/pricing", middlewares.Chain(controllers.PricingIndex, middlewares.Method("GET"), middlewares.Logging())).Methods("GET")

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
