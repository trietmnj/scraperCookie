package scraper

import (
	"github.com/gocolly/colly"
	"github.com/trietmnj/scraperCookie/rest"
	"github.com/trietmnj/scraperCookie/store"
)

type IScraper interface {
	Scrape() error // start scraping
	AddConfig(c CollectorConfig)
	AddStoreAccessor() // data storage accessor
	AddHandler(h CallbackHandler) error
}

type BaseScraper struct {
	list         rest.URLList
	FlagUseProxy bool // TODO future extension
	Store        store.StoreAccessor
	config       CollectorConfig
	handlers     []CallbackHandler
}

type CollectorConfig []func(*colly.Collector)

//
type CallbackHandler struct {
	order    string      // request, error, responseheader, response, html, xml, scraped
	optParam string      // optional param
	handler  interface{} // handler function http://go-colly.org/docs/introduction/start/
}

// Usage:
// s := EndpointScraper{}
// s.AddConfig([colly.Async(true)])
// s.AddURLs([]string{"https://httpbin.org/get"})
// s.AddStoreAccessor()
// s.Scrape()
type EndpointScraper struct {
	BaseScraper
}
