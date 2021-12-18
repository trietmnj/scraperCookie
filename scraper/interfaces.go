package scraper

import (
	"github.com/gocolly/colly"
	"github.com/trietmnj/scraperCookie/rest"
	"github.com/trietmnj/scraperCookie/store"
)

type IScraper interface {
	Scrape() error // start scraping
	AddConfig(c CollectorConfig)
	AddStoreAccessor(s store.IStore) // data storage accessor
	AddHandler(h CallbackHandler) error
}

type BaseScraper struct {
	FlagUseProxy bool // TODO future extension
	Store        store.IStore
	handlers     []CallbackHandler
}

type CollectorConfig []func(*colly.Collector)

type CallbackHandler struct {
	order string // request, error, responseheader, response, html, xml, scraped
	// optional param used with html or xml order
	// http://go-colly.org/docs/introduction/start/
	optParam string
	handler  interface{} // handler function used with any of the OnXXX methods
}

// implementation of of the base scraper
type EndpointScraper struct {
	BaseScraper
	list   rest.URLList
	config CollectorConfig
}
