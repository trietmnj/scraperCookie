package store

import (
	"fmt"
	"io"

	"github.com/trietmnj/scraperCookie/internal/types"
	"github.com/trietmnj/scraperCookie/pkg/config"
)

// Base interface, should not be fed directly to scraper
type IStore interface {
	init(c interface{}) error
	Storer
	Reader
}

type Storer interface {
	Store(l iLocator, data io.Reader) error // save into store
}

type Reader interface {
	Read(l iLocator) ([]byte, error)    // read data file
	KeyExists(l iLocator) (bool, error) // check if key is valid
	List(l iLocator) ([]Locator, error) // list of files
}

// Factory method to generate store
func NewStore(c config.StoreConfig) (IStore, error) {
	var s IStore
	switch c.StoreType {
	case types.S3Store:
		s = &s3Store{}
		s.init(c.S3Store)
	case types.LocalStore:
		s = &LocalStore{}
		s.init(c.LocalStore)
	default:
		return nil, fmt.Errorf("store factory: unable to generate new store")
	}
	return s, nil
}
