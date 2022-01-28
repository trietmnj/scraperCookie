package proxy

import (
	"errors"
	"path/filepath"
	"reflect"
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
		return nil, err
	}

	// path := os.Join(files[0].Bucket + "/" + files[0].Key

	var s5Slice []string
	switch l.Key {
	case "https://www.us-proxy.org/":

		// check if store is local to ensure StorePath field is available
		emptyLocalStore, err := store.NewStore("local")
		if err != nil {
			return nil, err
		}
		if reflect.TypeOf(s) == reflect.TypeOf(emptyLocalStore) {
			s, ok := s.(*store.LocalStore)
			if !ok {
				return nil, errors.New("unable to convert IStore to LocalStore")
			}
			path := filepath.Join(s.StorePath, files[len(files)-1].Bucket, files[len(files)-1].Key)
			d2Slice, err := utils.ReadCsv(path, true)
			if err != nil {
				return nil, err
			}
			d := []usProxy{}
			err = unmarshal(d2Slice, &d)
			if err != nil {
				return nil, err
			}

			for _, row := range d {
				s5Slice = append(s5Slice, "socks5://"+row.IP+":"+row.Port)
			}
		} else {
			return nil, errors.New("NewProxyFunction: unable to determine store type")
		}
	default:
	}

	return proxy.RoundRobinProxySwitcher(s5Slice...)
}
