package scraper

import (
	"errors"
	"fmt"

	"github.com/gocolly/colly"
	"github.com/trietmnj/scraperCookie/rest"
)

type Downloader struct {
	Collector colly.Collector
}

type RequestConfig struct {
	Endpoint  string
	Type      rest.RequestType // eg GET, POST, etc.
	URLParams map[string]string
}

func (p RequestConfig) AddEndPoint(s string) {
	p.Endpoint = s
}

// Send HTTP request and automatically parse any json response
func (d Downloader) SendRequest(
	rc RequestConfig,
	handleResponse func(r *colly.Response),
) (colly.Response, error) {

	c := colly.NewCollector(
		colly.AllowedDomains(d.ApiHost),
	)

	c.OnRequest(func(r *colly.Request) {
		r.Ctx.Put("url", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		// if r.StatusCode < 200 || r.StatusCode > 299 {
		// 	return nil, errors.New()
		// }

		// fmt.Println("Read from: " + r.Ctx.Get("url"))
		// return r, nil
		handleResponse(r)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	// switch requestType := p.Type; requestType {
	// case "GET":
	// 	url := d.ApiHost + p.Endpoint
	// 	fmt.Println("GET request: " + url)

	// 	response, err := http.Get(url)
	// 	defer response.Body.Close()
	// 	return response, err
	// default:
	// 	fmt.Println("Wrong request type")
	// }

	// if response.StatusCode == 200 {
	// 	bodyText, err := ioutil.ReadAll(response.Body)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	return bodyText, nil
	// } else {
	// 	return nil, errors.New("Request failed")
	// }
	return nil, errors.New("Failed to SendRequest")
}
