package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello")
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Hello)

	http.ListenAndServe(":8080", r)
}
