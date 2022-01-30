// Concrete builder
package scraper

import (
	"errors"

	"github.com/gocolly/colly"
	"github.com/trietmnj/scraperCookie/pkg/store"
)

type endpointBuilder struct {
	configs     []func(*colly.Collector)
	store       store.IStore
	handlers    []ResponseHandler
	selectors   []string
	urls        []string
	proxySwitch colly.ProxyFunc
}

func newEndPointScraperBuilder() *endpointBuilder {
	return &endpointBuilder{}
}

func (b *endpointBuilder) setConfig(c CollectorConfig) {
	b.configs = append(b.configs, c)
}

func (b *endpointBuilder) setHandler(h ResponseHandler) error {

	var w interface{}
	var ok bool
	switch h.order {
	case "error":
		w, ok = h.handler.(func(_ *colly.Response, err error))
	case "reponse":
		w, ok = h.handler.(func(r *colly.Response))
	default:
		ok = false
	}

	if !ok {
		return errors.New("Unable to assert handler into colly API defined callbacks")
	} else {
		b.handlers = append(b.handlers, ResponseHandler{h.order, h.optParam, w})
	}
	return nil
}

func (b *endpointBuilder) setStore(s store.IStore) {
	b.store = s
}

func (b *endpointBuilder) setSelectors(s []string) {
	b.selectors = s
}

func (b *endpointBuilder) setUrls(u []string) {
	b.urls = u
}

func (b *endpointBuilder) setProxySwitcher(p colly.ProxyFunc) {
	b.proxySwitch = p
}

func (b *endpointBuilder) getScraper() scraper {
	return scraper{
		configs:  b.configs,
		store:    b.store,
		handlers: b.handlers,
	}
}
