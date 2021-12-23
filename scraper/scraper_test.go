package scraper

import (
	"testing"

	"github.com/trietmnj/scraperCookie/store"
)

func Test_endpointScraper(t *testing.T) {
	endpointJsonBuilder := getScraperBuilder("EndpointJson")
	director := newDirector(endpointJsonBuilder)
	urls := []string{
		"https://httpbin.org/get",
		"http://localhost:3031/nsisapi/version",
	}
	s := store.S3JsonStore{}
	s.Init()
	endpointJsonScraper := director.buildScraper(&s)
	endpointJsonScraper.Scrape(urls)
}
