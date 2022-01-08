// Director
package scraper

import (
	"bytes"
	"fmt"
	"log"
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

// urlSelectors is a list of url and goquery selectors for htmlTableBuilder
func (d *director) BuildScraper(configJson string, s store.IStore, urlSelectors []string) (scraper, error) {
	var err error
	switch d.builder.(type) {

	case *endpointBuilder:
		d.builder.setConfig(colly.UserAgent(
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.93 Safari/537.36",
		))
		d.builder.setConfig(colly.Debugger(&debug.LogDebugger{}))
		d.builder.setStore(s)

		cfg, err := config.NewConfig(configJson)
		if err != nil {
			return scraper{}, err
		}

		// response handler
		d.builder.setHandler(ResponseHandler{
			"reponse", "",
			func(r *colly.Response) {
				if r.StatusCode == 200 {
					l := store.Locator{
						Key:    "ingest/" + cfg.Repo + "/" + strings.ReplaceAll(r.Request.URL.String(), "/", "-"),
						Bucket: cfg.Bucket,
					}
					s.Store(l, bytes.NewReader(r.Body))
				}
			},
		})

		// error handler
		d.builder.setHandler(ResponseHandler{
			"error", "",
			func(r *colly.Response, err error) {
				log.Fatal(err.Error())
			},
		})

	// htmlTableBuilder requires a CSV store
	case *htmlTableBuilder:
		d.builder.setConfig(colly.UserAgent(
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.93 Safari/537.36",
		))
		d.builder.setConfig(colly.Debugger(&debug.LogDebugger{}))
		d.builder.setStore(s)

		// urls slice not a multiple of 2
		if len(urlSelectors)%2 != 0 {
			return d.builder.getScraper(), fmt.Errorf("director: urls length must be even")
		}
		var urls []string
		var selectors []string
		for i, str := range urlSelectors {
			if i%2 == 0 {
				urls = append(urls, str)
			} else {
				selectors = append(selectors, str)
			}
		}
		d.builder.setSelectors(selectors)
		d.builder.setUrls(urls)

		// add handler specific to each url
		for i := 0; i < int(len(urls)/2); i++ {
			// html handler requires args from url
			d.builder.setHandler(ResponseHandler{
				order:    "html",
				optParam: urls[i*2+1],
				handler: func(e *colly.HTMLElement) {

					// doc, err := goquery.NewDocumentFromReader(strings.NewReader(e.Text))
					// if err != nil {
					// 	fmt.Errorf("response handler error: unable to parse html to table")
					// }

					// parse into 2d string matrix
					var data [][]string
					// iterate over rows
					e.ForEach("tr", func(idx int, e2 *colly.HTMLElement) {
						e2.ChildAttr()
						data = append(data)
					})

					// doc.Find(".spy1x").Each()
					l := store.Locator{}
					s.Store(l)
				},
			})
		}

		// error handler
		d.builder.setHandler(ResponseHandler{
			"error", "",
			func(r *colly.Response, err error) {
				fmt.Errorf("request error on url: " + r.Request.URL.String())
			},
		})

	default:
		err = fmt.Errorf("director: unable to construct scraper builder")
	}
	return d.builder.getScraper(), err
}
