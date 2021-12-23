// Concrete product
package scraper

import (
	"errors"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
	"github.com/trietmnj/scraperCookie/store"
)

type scraper struct {
	configs  []func(*colly.Collector) // configure colly collector
	store    store.IStore             // has Init(), Read(), and Write()
	handlers []ResponseHandler        // response handlers
}

func (s scraper) Scrape(urls []string) error {

	// add config
	c := colly.NewCollector(
		s.configs...,
	)

	// add URLs
	q, err := queue.New(
		5,
		&queue.InMemoryQueueStorage{MaxSize: 100000},
	)
	if err != nil {
		return err
	}
	for _, url := range urls {
		q.AddURL(url)
	}

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
			return errors.New("colly callback API not available for order: " + h.order)
		}
	}

	err = q.Run(c)
	if err != nil {
		return err
	}
	return nil
}
