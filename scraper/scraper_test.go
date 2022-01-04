package scraper

import (
	"github.com/trietmnj/scraperCookie/store"
)

func ExampleEndpointScraper() {
	endpointJsonBuilder := NewScraperBuilder("EndpointJson")
	director := NewDirector(endpointJsonBuilder)

	// urls is a slice of endpoints that returns json body
	urls := []string{
		"https://httpbin.org/get",
		"http://localhost:3031/nsisapi/version",
	}
	s := store.TextStore{}
	s.Init()
	endpointJsonScraper := director.BuildScraper(&s, urls)
	endpointJsonScraper.Scrape()
}

// ExampleHtmlTableScraper is an example of a scraper that parse
// data based on an HTML <table> tag
func ExampleHtmlTableScraper() {
	// Director directs specific builder
	htmlTableBuilder := NewScraperBuilder("HtmlTable")
	director := NewDirector(htmlTableBuilder)

	// first string contains url,
	// second string contains table selector
	// third string contains another url
	// fourth string contains the selector for the url in the third string
	// each url should be coupled with a selector
	urlHtml := []string{
		"https://spys.one/en/socks-proxy-list/",
		"td table:first",
		// "https://www.us-proxy.org/",
		// "table.table-responsive.fpl-list",
	}

	s := store.NewS3JsonStore()
	htmlTableScraper := director.BuildScraper(s, urlHtml)
	htmlTableScraper.Scrape()
}
