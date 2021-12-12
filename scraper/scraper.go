package scraper

import (
	"errors"

	"github.com/gocolly/colly"
)

// Common methods accessible via the BaseScraper

func (s *BaseScraper) AddStoreAccessor() {

}

func (s *BaseScraper) AddHandler(h CallbackHandler) error {

	// validate that handler can fit into colly callback API
	var w interface{}
	var ok bool
	switch h.order {
	case "request":
		w, ok = h.handler.(func(r *colly.Request))
	case "error":
		w, ok = h.handler.(func(_ *colly.Response, err error))
	// case "reponse-headers": // TODO update colly to a version with c.OnResponseHeaders()
	// 	w, ok = h.handler.(func(r *colly.Response))
	case "reponse":
		w, ok = h.handler.(func(r *colly.Response))
	case "html":
		w, ok = h.handler.(func(e *colly.HTMLElement))
	case "xml":
		w, ok = h.handler.(func(e *colly.XMLElement))
	case "scraped":
		w, ok = h.handler.(func(r *colly.Response))
	default:
		ok = false
	}

	if !ok {
		return errors.New("Unable to assert handler into colly API defined callbacks")
	} else {
		s.handlers = append(s.handlers, CallbackHandler{h.order, h.optParam, w})
		return nil
	}
}
