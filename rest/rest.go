package rest

import "github.com/trietmnj/scraperCookie/utils"

type RequestType utils.Bits

const (
	GET RequestType = 1 << iota
	POST
	DELETE
	PUT
	UPDATE
)

type WebPage struct {
	Url     string `json:"url"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type RequestConfig struct {
	Endpoint  string
	Type      RequestType // eg GET, POST, etc.
	URLParams map[string]string
}
