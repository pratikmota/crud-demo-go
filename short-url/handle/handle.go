package handle

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// POST
func ShortUrl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req ShortRequest
	_ = json.NewDecoder(r.Body).Decode(&req)
	json.NewEncoder(w).Encode(&req)
}

// GET
func RedirectURL(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello")
	//
	//https://dev.to/envitab/how-to-build-a-url-shortener-with-go-5hn5
}
