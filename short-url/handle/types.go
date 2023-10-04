package handle

type ShortRequest struct {
	LongUrl         string `json:"url"`
	MaxRequestCount int    `json:"max_count"`
	IPAddress       string `json:"ip"`
}
