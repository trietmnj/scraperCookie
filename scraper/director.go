// Director
package scraper

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"github.com/trietmnj/scraperCookie/store"
)

type director struct {
	builder iBuilder
}

func newDirector(b iBuilder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) setBuilder(b iBuilder) {
	d.builder = b
}

func (d *director) buildScraper(s store.IStore) scraper {
	d.builder.setConfig(colly.UserAgent(
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.93 Safari/537.36",
	))
	d.builder.setConfig(colly.Debugger(&debug.LogDebugger{}))
	d.builder.setStore(s)
	d.builder.setHandler(ResponseHandler{
		"reponse", "",
		func(r *colly.Response) {
			fmt.Println("Visited", r.Request.URL)
			if r.StatusCode == 200 {
				l := store.Locator{
					cfg.Bucket, cfg.DataSource, cfg.RepoName, strings.ReplaceAll(r.Request.URL.String(), "/", "-"),
				}
				s.Store(l, bytes.NewReader(r.Body))
			}
		},
	})
	return d.builder.getScraper()
}
