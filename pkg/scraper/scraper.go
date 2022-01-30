// Concrete product
package scraper

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
	"github.com/trietmnj/scraperCookie/pkg/store"
)

type scraper struct {
	configs     []func(*colly.Collector) // configure colly collector
	store       store.IStore             // has Init(), Read(), and Write()
	handlers    []ResponseHandler        // response handlers
	urls        []string                 // list of urls to query
	selectors   []string                 // optional string slice with selectors specific to each url
	proxySwitch colly.ProxyFunc
}

func (s scraper) Scrape() error {

	// config
	c := colly.NewCollector(
		s.configs...,
	)

	// URLs
	q, err := queue.New(
		5,
		&queue.InMemoryQueueStorage{MaxSize: 100000},
	)
	if err != nil {
		return err
	}
	for _, url := range s.urls {
		q.AddURL(url)
	}

	// round robin proxy switcher
	if s.proxySwitch != nil {
		c.SetProxyFunc(s.proxySwitch)
	}

	// handlers
	for _, h := range s.handlers {
		switch h.order {
		case "request":
			c.OnRequest(h.handler.(func(r *colly.Request)))
		case "error":
			c.OnError(h.handler.(func(_ *colly.Response, err error)))
		case "reponse":
			c.OnResponse(h.handler.(func(r *colly.Response)))
		case "html":
			c.OnHTML(h.optParam, h.handler.(func(e *colly.HTMLElement)))
		case "xml":
			c.OnXML(h.optParam, h.handler.(func(e *colly.XMLElement)))
		case "scraped":
			c.OnScraped(h.handler.(func(r *colly.Response)))
		default:
			return fmt.Errorf("scraper: invalid handler order - " + h.order)
		}
	}

	err = q.Run(c)
	if err != nil {
		return err
	}
	return nil
}
