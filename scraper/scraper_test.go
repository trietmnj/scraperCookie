package scraper

import (
	"fmt"
	"testing"

	"github.com/gocolly/colly"
	"github.com/stretchr/testify/assert"
	"github.com/trietmnj/scraperCookie/rest"
)

func TestScrape(t *testing.T) {

	c := rest.RequestConfig{
		Endpoint:  "/get",
		Type:      rest.GET,
		URLParams: map[string]string{}, // empty map
	}

	s := Scraper{
		Collector:    colly.NewCollector(),
		FlagUseProxy: false,
	}

	s.AddDomain("https://httpbin.org")

	s.Scrape(c, func(r *colly.Response) {
		fmt.Println(r.StatusCode)
		assert.Equal(t, r.StatusCode, 200)
	})
}
