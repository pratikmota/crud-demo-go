package handle

import (
	"fmt"
	"net/http"
)

func ShortUrl(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello")
}
func RedirectURL(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello")
}
