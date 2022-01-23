package proxy

import (
	"errors"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
	"github.com/trietmnj/scraperCookie/store"
	"github.com/trietmnj/scraperCookie/utils"
)

// NewProxyFunction generates a proxy function from a store - key in Locator should be source ie site url
func NewProxyFunction(s store.IStore, l store.Locator) (colly.ProxyFunc, error) {
	key := "ingest/proxy/" + strings.ReplaceAll(l.Key, "/", "%2F")
	files, err := s.List(store.Locator{
		Key:    key,
		Bucket: l.Bucket,
	})
	if err != nil {
		return nil, errors.New("NewProxyFunction: unable to generate list from store")
	}

	path := files[0].Bucket + files[0].Key

	var s5Slice []string
	switch l.Key {
	case "https://www.us-proxy.org/":
		d := usProxyData{}
		d2Slice, err := utils.ReadCsv(path, false)
		if err != nil {
			return nil, errors.New("NewProxyFunction: unable to ")
		}
		d.Marshal(d2Slice)

		for _, row := range d {
			s5Slice = append(s5Slice, "socks5://"+row.IP+":"+row.Port)
		}

	default:
	}

	return proxy.RoundRobinProxySwitcher(s5Slice...)
}
