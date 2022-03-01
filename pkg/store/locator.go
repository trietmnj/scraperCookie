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
	path string
	file string
}

type S3Location struct {
	bucket string
	key    string
}

type Locator struct {
	storeType types.Store
	local     LocalLocation
	s3        S3Location
}

func (l Locator) Container() string {
	switch l.storeType {
	case types.LocalStore:
		return strings.ReplaceAll(l.local.path, l.local.file, "")
	case types.S3Store:
		return l.s3.bucket
	}
	return ""
}

func (l Locator) Path() string {
	switch l.storeType {
	case types.LocalStore:
		return l.Container()
	case types.S3Store:
		return l.File()
	}
	return ""
}

func (l Locator) File() string {
	switch l.storeType {
	case types.LocalStore:
		return l.local.file
	case types.S3Store:
		return l.s3.key
	}
	return ""
}
