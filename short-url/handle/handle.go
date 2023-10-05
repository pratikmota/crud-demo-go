package handle

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

// POST
func ShortUrl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req ShortRequest
	reqBody, _ := io.ReadAll(r.Body)
	json.Unmarshal(reqBody, &req)

	//json.NewEncoder(w).Encode(req)

	// LOGIC for SHORT URL

	response := ShortResponse{
		LongUrl:  req.LongUrl,
		ShortUrl: "test.com/xyz",
	}
	responseJson, err := json.Marshal(response)
	if err != nil {
		panic(fmt.Errorf("error while marshaling UserDetailsResponse response: %s", err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)

}

// GET
func RedirectURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	shortcode := vars["shortcode"]

	// find shortcode from database and respective URL

	url := fmt.Sprintf("text.com/%s", shortcode)
	resp := ShortRequest{
		LongUrl: url,
	}

	rjson, _ := json.Marshal(resp)
	w.WriteHeader(http.StatusOK)
	w.Write(rjson)

	//https://dev.to/envitab/how-to-build-a-url-shortener-with-go-5hn5
}
