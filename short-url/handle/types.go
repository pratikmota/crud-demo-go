package handle

type ShortRequest struct {
	LongUrl string `json:"url"`
}

type ShortResponse struct {
	LongUrl  string `json:"url"`
	ShortUrl string `json:"shorturl"`
}
