package scraper

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trietmnj/scraperCookie/config"
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
	c, _ := config.NewConfig("data/config.json")
	endpointJsonScraper, _ := director.BuildScraper(c, s, urls)
	endpointJsonScraper.Scrape()
}

// ExampleHtmlTableScraper is an example of a scraper that parse
// data based on an HTML <table> tag
func TestHtmlTableScraper(t *testing.T) {
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
		"table",
		// "td table:first",
		// "https://www.us-proxy.org/",
		// "table.table-responsive.fpl-list",
	}
	s, err := store.NewStore("s3")
	assert.Nil(t, err)
	c, err := config.NewConfig("/workspaces/scraperCookie/data/config.json")
	assert.Nil(t, err)
	htmlTableScraper, err := director.BuildScraper(c, s, urlHtml)
	assert.Nil(t, err)
	htmlTableScraper.Scrape()
}
