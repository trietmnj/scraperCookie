// Builder interface
package scraper

import (
	"github.com/gocolly/colly"
	"github.com/trietmnj/scraperCookie/internal/types"
	"github.com/trietmnj/scraperCookie/pkg/store"
)

type iBuilder interface {
	setConfig(c CollectorConfig)
	setStore(s store.IStore) // data storage accessor
	setHandler(h ResponseHandler) error
	setSelectors(s []string)
	setUrls(u []string)
	setProxySwitcher(p colly.ProxyFunc)
	getScraper() scraper
}

func NewScraperBuilder(dataType types.Data) iBuilder {
	switch dataType {
	case types.JSONEndpointData:
		return &endpointBuilder{}
	case types.HTMLTableData:
		return &htmlTableBuilder{}
	default:
		return nil
	}
}

type CollectorConfig func(*colly.Collector)

type ResponseHandler struct {
	order    string      // request, error, responseheader, response, html, xml, scraped
	optParam string      // optional param used with OnHtml or OnXML
	handler  interface{} // handler function used with any of the OnXXX methods
}
