package scraper

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/trietmnj/scraperCookie/rest"
)

type Scraper struct {
	Collector    *colly.Collector
	FlagUseProxy bool // TODO future extension
}

// Add another domain to collector
func (s Scraper) AddDomain(n string) {
	s.Collector.AllowedDomains = append(s.Collector.AllowedDomains, n)
}

// Start scraping
func (s Scraper) Scrape(
	rc rest.RequestConfig,
	handleResponse func(r *colly.Response),
) {

	s.Collector.OnRequest(func(r *colly.Request) {
		r.Ctx.Put("url", r.URL.String())
	})

	s.Collector.OnResponse(func(r *colly.Response) {
		handleResponse(r)
	})

	s.Collector.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})
}
