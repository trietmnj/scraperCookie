package store

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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
			Bucket: "repo",
			Key:    "test.json",
		},
		strings.NewReader(
			// `{"title":"Survey Test","description":"This is a description of the test survey","active":true}`,
			`{"title":"Survey Test","description":"This is a description of the test survey","active":false}`,
		))
	fmt.Println(err)
}

func TestLocalStoreList(t *testing.T) {
	s, err := NewStore("local")
	source := "https://www.us-proxy.org/"

	// Locator key is source url for List method
	l := Locator{
		Bucket: "finance-lake",
		Key:    filepath.Join("ingest/proxy/", strings.ReplaceAll(source, "/", "%2F")),
	}

	files, err := s.List(l)
	assert.Nil(t, err, "store List() has error")

	for _, file := range files {
		exists, err := s.KeyExists(file)
		assert.Nil(t, err, "error in KeyExists for key: "+file.Key)
		assert.True(t, exists, "KeyExists should be true for key: "+file.Key)
	}
}

func TestLocalStoreRead(t *testing.T) {

	s, err := NewStore("local")
	source := "https://www.us-proxy.org/"

	// Locator key is source url for List method
	l := Locator{
		Bucket: "finance-lake",
		Key:    filepath.Join("ingest/proxy/", strings.ReplaceAll(source, "/", "%2F"), ""),
	}

	locators, err := s.List(l)
	assert.Nil(t, err, "store List() has error")

	for _, l := range locators {
		file, err := s.Read(l)
		assert.Nil(t, err, "error in Read for key: "+file.Key)
		fmt.Println(file)
	}
}
