// TODO Dangling dependency - consider removing
package rest

import "github.com/trietmnj/scraperCookie/utils"

type RequestType utils.Bits

type WebPage struct {
	Url     string `json:"url"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type RequestConfig struct {
	Endpoint  string
	Type      string // eg GET, POST, etc.
	URLParams map[string]string
}

// List of URLs to query
type URLList []string
