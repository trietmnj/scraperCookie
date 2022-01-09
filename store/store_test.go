package store

import (
	"fmt"
	"strings"
	"testing"
)

func ExampleS3StoreUpload(t *testing.T) {
	s, err := NewStore("s3")
	s.Init()
	err = s.Store(
		Locator{"finance-lake", "bronze/ingest/vic/ideasum-json/test.json"},
		strings.NewReader(
			`{"title":"Survey Test","description":"This is a description of the test survey","active":true}`,
		))
	fmt.Println(err)
}

// // TODO add in Locator fields correspond to new structures
// func Test_keyExists(t *testing.T) {
// 	s.Init()
// 	exists, _ := s.KeyExists(Locator{"finance-lake", "bronze/ingest/vic/ideasum-json/test.json"})
// 	fmt.Println(exists)
// 	exists, _ = s.KeyExists(Locator{"finance-lake", "bronze/ingest/vic/ideasum-json/test"})
// 	fmt.Println(exists)
// }

// func ExampleS3Store() {
// 	s, err := NewStore("s3")
// 	if err != nil {
// 		panic("unable initialize s3 store")
// 	}
// 	l := Locator{"finance-lake"}
// 	s.Store(l, strings.NewReader(""))

// }

func TestLocalStore(t *testing.T) {
	s, err := NewStore("local")
	err = s.Store(
		Locator{
			Bucket: "",
			Key:    "/workspaces/cookieScraper/data/test.json",
		},
		strings.NewReader(
			`{"title":"Survey Test","description":"This is a description of the test survey","active":true}`,
		))
	fmt.Println(err)
}
