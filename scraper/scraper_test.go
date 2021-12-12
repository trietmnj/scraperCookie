package scraper

import (
	"fmt"
	"log"
	"testing"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
)

func Test_Scrape(t *testing.T) {

	s := EndpointScraper{}

	s.AddConfig(
		[]func(*colly.Collector){
			colly.UserAgent(
				"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.93 Safari/537.36",
			),
			colly.Debugger(&debug.LogDebugger{}),
			// TODO figure out async stuff
			// With the EndpointScraper, queue ends before any of the async stuff
			//  are completed, unable to validate if request is successful
			// colly.Async(true),
		},
	)
	s.AddURLs(
		[]string{
			"https://httpbin.org/get",
			"http://localhost:3031/nsisapi/version",
		},
	)
	s.AddHandler(CallbackHandler{
		"request", "",
		func(r *colly.Request) {
			fmt.Println("Visiting", r.URL)
		},
	})
	s.AddHandler(CallbackHandler{
		"reponse", "",
		func(r *colly.Response) {
			fmt.Println("Visited", r.Request.URL)
			fmt.Println("StatusCode: ", r.StatusCode)
		},
	})
	s.AddHandler(CallbackHandler{"scraped", "",
		func(r *colly.Response) {
			fmt.Println("Finished", r.Request.URL)
		},
	})
	s.AddHandler(CallbackHandler{"error", "",
		func(_ *colly.Response, err error) {
			log.Println("Something went wrong:", err)
		},
	})
	// s.AddStoreAccessor()
	s.Scrape()
}
