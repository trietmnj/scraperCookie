package scraper

import (
	"errors"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
	"github.com/trietmnj/scraperCookie/rest"
)

func (s EndpointScraper) Scrape() error {

	// add config
	c := colly.NewCollector(
		s.config...,
	)

	// add URLs
	q, err := queue.New(
		2,
		&queue.InMemoryQueueStorage{MaxSize: 100000},
	)
	if err != nil {
		return err
	}
	for _, url := range s.list {
		q.AddURL(url)
	}

	// add handlers
	// TODO find a safer approach
	// TODO update colly to a version with colly.Collector.OnResponseHeaders available
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
			c.OnXML(h.optParam, h.handler.(func(e *colly.HTMLElement)))
		case "scraped":
			c.OnResponse(h.handler.(func(r *colly.Response)))
		default:
			errors.New("colly callback API not available for order: " + h.order)
		}
	}

	q.Run(c)
	return nil
}

// Add a list of URLs to the scraper. Unique for EndpointScraper.
func (s *EndpointScraper) AddURLs(l rest.URLList) {
	s.list = l
}

// Add optional *colly.Collector configs
func (s *EndpointScraper) AddConfig(c CollectorConfig) {

	requiredConfigs := []func(*colly.Collector){
		colly.MaxDepth(1),
		colly.Async(true),
	}

	s.config = append(c, requiredConfigs...)
}
