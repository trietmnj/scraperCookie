// Concrete builder
package scraper

import (
	"errors"

	"github.com/gocolly/colly"
	"github.com/trietmnj/scraperCookie/pkg/store"
)

type htmlTableBuilder struct {
	configs     []func(*colly.Collector)
	store       store.IStore
	handlers    []ResponseHandler
	urls        []string
	selectors   []string
	proxySwitch colly.ProxyFunc
}

func newHtmlTableBuilder() *htmlTableBuilder {
	return &htmlTableBuilder{}
}

func (b *htmlTableBuilder) setConfig(c CollectorConfig) {
	b.configs = append(b.configs, c)
}

func (b *htmlTableBuilder) setHandler(h ResponseHandler) error {

	var w interface{}
	var ok bool
	switch h.order {
	case "error":
		w, ok = h.handler.(func(_ *colly.Response, err error))
	case "html":
		w, ok = h.handler.(func(e *colly.HTMLElement))
		if h.optParam == "" {
			ok = false
		}
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

func (b *htmlTableBuilder) setStore(s store.IStore) {
	b.store = s
}

func (b *htmlTableBuilder) setSelectors(s []string) {
	b.selectors = s
}

func (b *htmlTableBuilder) setUrls(u []string) {
	b.urls = u
}

func (b *htmlTableBuilder) setProxySwitcher(p colly.ProxyFunc) {
	b.proxySwitch = p
}

func (b *htmlTableBuilder) getScraper() scraper {
	return scraper{
		configs:     b.configs,
		store:       b.store,
		handlers:    b.handlers,
		urls:        b.urls,
		selectors:   b.selectors,
		proxySwitch: b.proxySwitch,
	}
}
