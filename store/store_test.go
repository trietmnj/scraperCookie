package store

import (
	"fmt"
	"strings"
	"testing"
)

var s = S3JsonStore{}

func Test_upload(t *testing.T) {
	s.Init()
	err := s.Store(
		Locator{"finance-lake", "bronze/ingest/vic/ideasum-json/test.json"},
		strings.NewReader(
			`{"title":"Survey Test","description":"This is a description of the test survey","active":true}`,
		))
	fmt.Println(err)
}

func Test_keyExists(t *testing.T) {
	s.Init()
	exists, _ := s.KeyExists(Locator{"finance-lake", "bronze/ingest/vic/ideasum-json/test.json"})
	fmt.Println(exists)
	exists, _ = s.KeyExists(Locator{"finance-lake", "bronze/ingest/vic/ideasum-json/test"})
	fmt.Println(exists)
}
