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
