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
	s, _ := store.NewStore("s3")
	endpointJsonScraper, _ := director.BuildScraper("data/config.json", s, urls)
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
	s, _ := store.NewStore("s3")
	htmlTableScraper, _ := director.BuildScraper("data/config.json", s, urlHtml)
	htmlTableScraper.Scrape()
}
