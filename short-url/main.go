package main

import (
	"net/http"

	"short-url/handle"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/short", handle.ShortUrl).Methods("POST")
	r.HandleFunc("/{shortcode}", handle.RedirectURL).Methods("GET")

	http.ListenAndServe(":8080", r)
}
