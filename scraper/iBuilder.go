// Builder interface
package scraper

import (
	"github.com/gocolly/colly"
	"github.com/trietmnj/scraperCookie/store"
)

type iBuilder interface {
	setConfig(c CollectorConfig)
	setStore(s store.IStore) // data storage accessor
	setHandler(h ResponseHandler) error
	setSelectors(s []string)
	setUrls(u []string)
	getScraper() scraper
}

func NewScraperBuilder(scraperType string) iBuilder {
	switch scraperType {
	case "EndpointJson":
		return &endpointBuilder{}
	case "HtmlTable":
		return &htmlTableBuilder{}
	default:
		// errors.New("Unable to create scraper object for scraperType: " + scrapescraperType)
		return nil
	}
}

type CollectorConfig func(*colly.Collector)

type ResponseHandler struct {
	order    string      // request, error, responseheader, response, html, xml, scraped
	optParam string      // optional param used with OnHtml or OnXML
	handler  interface{} // handler function used with any of the OnXXX methods
}
