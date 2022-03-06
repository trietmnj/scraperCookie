package store

import (
	"strings"

	"github.com/trietmnj/scraperCookie/internal/types"
)

type iLocator interface {
	Container() string // Bucket for s3, same as path for local
	Path() string      // Path do not include file name
	File() string
}

type LocalLocation struct {
	Path string
	File string
}

type S3Location struct {
	Bucket string
	Key    string
}

type Locator struct {
	StoreType types.Store
	Local     LocalLocation
	S3        S3Location
}

func (l Locator) Container() string {
	switch l.StoreType {
	case types.LocalStore:
		return strings.ReplaceAll(l.Local.Path, l.Local.File, "")
	case types.S3Store:
		return l.S3.Bucket
	}
	return ""
}

func (l Locator) Path() string {
	switch l.StoreType {
	case types.LocalStore:
		return l.Container()
	case types.S3Store:
		return l.File()
	}
	return ""
}

func (l Locator) File() string {
	switch l.StoreType {
	case types.LocalStore:
		return l.Local.File
	case types.S3Store:
		return l.S3.Key
	}
	return ""
}
