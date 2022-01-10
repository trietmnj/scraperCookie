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
func (d *director) BuildScraper(c config.Config, s store.IStore, urlSelectors []string) (scraper, error) {

	// Common components for all scraper types
	d.builder.setConfig(colly.UserAgent(
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.93 Safari/537.36",
	))
	d.builder.setConfig(colly.Debugger(&debug.LogDebugger{}))
	d.builder.setStore(s)

	//TODO - add in date/time meta to folder structure

	var err error
	switch d.builder.(type) {

	case *endpointBuilder:
		if err != nil {
			return scraper{}, err
		}

		// response handler
		d.builder.setHandler(ResponseHandler{
			"reponse", "",
			func(r *colly.Response) {
				if r.StatusCode == 200 {
					l := store.Locator{
						Key:    "ingest/" + c.Repo + "/" + strings.ReplaceAll(r.Request.URL.String(), "/", "-"),
						Bucket: c.Bucket,
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
		for _, selector := range selectors {
			// html handler requires args from url
			d.builder.setHandler(ResponseHandler{
				order:    "html",
				optParam: selector,
				handler: func(table *colly.HTMLElement) {
					var data []string // parse into 2d string matrix
					var rowData []string
					table.ForEach("tr", func(_ int, row *colly.HTMLElement) {
						rowData = make([]string, 0)
						row.ForEach("td", func(_ int, cell *colly.HTMLElement) {
							rowData = append(rowData, cell.Text)
						})
						data = append(data, strings.Join(rowData, ","))
					})
					dataStr := strings.Join(data, "\n")

					// search for new key
					var key string
					key = "ingest/" + c.Repo + "/" + strings.ReplaceAll(table.Request.URL.String(), "/", "-") + "/" + "table-"
					var exists bool
					var i int
					i = 0
					exists = true
					for exists {
						exists, _ = s.KeyExists(store.Locator{Key: key + string(i), Bucket: c.Bucket})
					}

					s.Store(store.Locator{
						Key:    key + string(i),
						Bucket: c.Bucket,
					}, strings.NewReader(dataStr))
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
