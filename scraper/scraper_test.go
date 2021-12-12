package scraper

import (
	"fmt"
	"testing"

	"github.com/gocolly/colly"
)

func Test_Scrape(t *testing.T) {

	s := EndpointScraper{}

	s.AddConfig(
		[]func(*colly.Collector){colly.Async(true)},
	)
	s.AddURLs(
		[]string{"https://httpbin.org/get"},
	)
	s.AddHandler(CallbackHandler{
		"reponse", "",
		func(r *colly.Response) {
			fmt.Println("Visited", r.Request.URL)
		}})
	s.AddStoreAccessor()
	s.Scrape()
}
