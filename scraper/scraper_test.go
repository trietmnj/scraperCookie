package scraper

import (
	"testing"

	"github.com/trietmnj/scraperCookie/store"
)

func Test_endpointScraper(t *testing.T) {
	endpointJsonBuilder := GetScraperBuilder("EndpointJson")
	director := NewDirector(endpointJsonBuilder)
	urls := []string{
		"https://httpbin.org/get",
		"http://localhost:3031/nsisapi/version",
	}
	s := store.S3JsonStore{}
	s.Init()
	endpointJsonScraper := director.BuildScraper(&s, urls)
	endpointJsonScraper.Scrape()
}

func ExampleHtmlTableScraper() {
	htmlTableBuilder := NewScraperBuilder("HtmlTable")
	director := NewDirector(htmlTableBuilder)
	urls := []string{
		"https://spys.one/en/socks-proxy-list/",
		"td table:first",
		"https://www.us-proxy.org/",
		"table.table-responsive.fpl-list",
	}
	s := store.NewS3JsonStore()
	htmlTableScraoer := director.BuildScraper(s, urls)
	htmlTableScraoer.Scrape()
}
