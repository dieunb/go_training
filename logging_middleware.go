package main

import (
	"log"
	"net/http"
)

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("request path " + r.URL.Path)
		f(w, r)
	}
}
