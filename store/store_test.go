package store

import (
	"fmt"
	"strings"
	"testing"
)

var s = TextStore{}

func Test_upload(t *testing.T) {
	s.Init()
	err := s.Store(
		Locator{"finance-lake", "bronze/ingest/vic/ideasum-json/test.json"},
		strings.NewReader(
			`{"title":"Survey Test","description":"This is a description of the test survey","active":true}`,
		))
	fmt.Println(err)
}

// TODO add in Locator fields correspond to new structures
func Test_keyExists(t *testing.T) {
	s.Init()
	exists, _ := s.KeyExists(Locator{"finance-lake", "bronze/ingest/vic/ideasum-json/test.json"})
	fmt.Println(exists)
	exists, _ = s.KeyExists(Locator{"finance-lake", "bronze/ingest/vic/ideasum-json/test"})
	fmt.Println(exists)
}

func ExampleTextStore() {
	s, err := NewStore("text")
	if err != nil {
	}

	s := &CSVStore{}
	l := Locator{"finance-lake"}
	s.Store(l, strings.NewReader(""))

}
