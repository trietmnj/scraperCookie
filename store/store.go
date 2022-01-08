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
}

// Factory method to generate store
func NewStore(sType string) (IStore, error) {
	switch sType {
	case "s3":
		s := S3Store{}
		s.Init()
		return &s, nil
	default:
		return &S3Store{}, fmt.Errorf("store: unable to generate new store")
	}
}
