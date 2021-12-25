// Director
package scraper

import (
	"bytes"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"github.com/trietmnj/scraperCookie/config"
	"github.com/trietmnj/scraperCookie/store"
)

type director struct {
	builder iBuilder
}

func NewDirector(b iBuilder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) setBuilder(b iBuilder) {
	d.builder = b
}

// TODO refactor ResponseHandler
// current ResponseHandler requires env cfg dependencies:
//  Bucket
//  DataSource
//  RepoName
func (d *director) BuildScraper(s store.IStore) scraper {
	d.builder.setConfig(colly.UserAgent(
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.93 Safari/537.36",
	))
	d.builder.setConfig(colly.Debugger(&debug.LogDebugger{}))
	d.builder.setStore(s)

	cfg := config.Init()
	d.builder.setHandler(ResponseHandler{
		"reponse", "",
		func(r *colly.Response) {
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
