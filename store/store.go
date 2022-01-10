package store

import (
	"fmt"
	"io"
)

// Locator or index to find data in store
type Locator struct {
	Key    string
	Bucket string
}

// Base interface, should not be fed directly to scraper
type IStore interface {
	Init()
	Store(l Locator, data io.Reader) error
	Read(l Locator) []byte
	KeyExists(l Locator) (bool, error)
}

// Factory method to generate store
func NewStore(sType string) (IStore, error) {
	var s IStore
	switch sType {
	case "s3":
		s = &s3Store{}
	case "local":
		s = &localStore{}
	default:
		return nil, fmt.Errorf("store factory: unable to generate new store")
	}
	s.Init()
	return s, nil
}
