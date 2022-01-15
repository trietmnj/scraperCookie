// Director
package scraper

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"time"

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

	// used to filter out tags that do not include data
	// invalidTags := []string{"script"}

	// Common components for all scraper types
	d.builder.setConfig(colly.UserAgent(
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.93 Safari/537.36",
	))
	d.builder.setConfig(colly.Debugger(&debug.LogDebugger{}))
	d.builder.setStore(s)

	dt := time.Now().UTC()
	year, month, date := dt.Date()
	hour := dt.Hour()
	min := dt.Minute()
	sec := dt.Second()

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

					key := "ingest/" + c.Repo + "/" +
						strings.ReplaceAll(r.Request.URL.String(), "/", "%2F") +
						fmt.Sprintf("%04d/%02d/%02d/%02d%02d%02d", year, int(month), date, hour, min, sec)

					if !strings.HasSuffix(key, ".json") {
						key += ".json"
					}

					l := store.Locator{
						Key:    key,
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
			return d.builder.getScraper(), fmt.Errorf("director: urlSelectors length must be even")
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

					// parse into 2d string matrix
					var data, rowData, headings, filtered []string

					// row
					table.ForEach("tr", func(_ int, row *colly.HTMLElement) {

						// heading
						row.ForEach("th", func(_ int, tableheading *colly.HTMLElement) {
							headings = append(headings, tableheading.Text)
						})

						// body
						row.ForEach("td", func(_ int, cell *colly.HTMLElement) {
							// cellData = make([]string, 0)
							// // jQueryNotFiler := ":not(" + strings.Join(invalidTags, ",") + ")" // filter out tags that do not include data
							// // cell.ForEach(jQueryNotFiler, func(_ int, e *colly.HTMLElement) {
							// cell.ForEach("*", func(_ int, e *colly.HTMLElement) {
							// 	cellData = append(cellData, e.Text)
							// })
							// rowData = append(rowData, strings.Join(cellData, ":"))
							rowData = append(rowData, cell.Text)
						})
						data = append(data, strings.Join(rowData, ","))
						rowData = nil
					})

					// filter out empty rows
					for _, val := range data {
						if val != "" {
							filtered = append(filtered, val)
						}
					}

					headingsSlice := []string{strings.Join(headings, ",")}
					dataStr := strings.Join(append(headingsSlice, filtered...), "\n")

					// TODO file management needs a bit more work
					// search for new key
					var key string
					// key using url encoding
					// https://www.w3schools.com/tags/ref_urlencode.ASP
					// page could have multiple tables with data
					key = "ingest/" + c.Repo + "/" +
						strings.ReplaceAll(table.Request.URL.String(), "/", "%2F") + "/" +
						fmt.Sprintf("%04d/%02d/%02d/%02d%02d%02d/table", year, int(month), date, hour, min, sec)
					var exists bool
					var i int
					i = 1
					exists = true
					var keyWithOrder string
					for exists {
						keyWithOrder = key + fmt.Sprintf("%02d.csv", i)
						exists, _ = s.KeyExists(store.Locator{Key: keyWithOrder, Bucket: c.Bucket})
						i++
					}

					// write only if there is data
					if len(dataStr) > 0 {
						s.Store(store.Locator{
							Key:    keyWithOrder,
							Bucket: c.Bucket,
						}, strings.NewReader(dataStr))
					}
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
