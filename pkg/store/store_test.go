package store

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trietmnj/scraperCookie/internal/types"
	"github.com/trietmnj/scraperCookie/pkg/config"
)

// func ExampleS3StoreUpload(t *testing.T) {
// 	s, err := NewStore("s3")
// 	s.Init()
// 	err = s.Store(
// 		Locator{"finance-lake", "bronze/ingest/vic/ideasum-json/test.json"},
// 		strings.NewReader(
// 			`{"title":"Survey Test","description":"This is a description of the test survey","active":true}`,
// 		))
// 	fmt.Println(err)
// }

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
	c, err := config.NewConfig(types.JsonConfigSource, "/workspaces/scraperCookie/test/config-local.json")
	assert.Nil(t, err)
	s, err := NewStore(c.StoreConfig)
	assert.Nil(t, err)
	err = s.Store(
		Locator{
			StoreType: types.LocalStore,
			Local: LocalLocation{
				Path: "repo",
				File: "test.json",
			},
		},
		strings.NewReader(
			`222Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do
            eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut
            enim ad minim veniam, quis nostrud exercitation ullamco laboris
            nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in
            reprehenderit in voluptate velit esse cillum dolore eu fugiat
            nulla pariatur. Excepteur sint occaecat cupidatat non proident,
            sunt in culpa qui officia deserunt mollit anim id est laborum.`,
		))
	assert.Nil(t, err)
}

// func TestLocalStoreList(t *testing.T) {
// 	s, err := NewStore("local")
// 	source := "https://www.us-proxy.org/"

// 	// Locator key is source url for List method
// 	l := Locator{
// 		Bucket: "finance-lake",
// 		Key:    filepath.Join("ingest/proxy/", strings.ReplaceAll(source, "/", "%2F")),
// 	}

// 	files, err := s.List(l)
// 	assert.Nil(t, err, "store List() has error")

// 	for _, file := range files {
// 		exists, err := s.KeyExists(file)
// 		assert.Nil(t, err, "error in KeyExists for key: "+file.Key)
// 		assert.True(t, exists, "KeyExists should be true for key: "+file.Key)
// 	}
// }

// func TestLocalStoreRead(t *testing.T) {

// 	s, err := NewStore("local")
// 	source := "https://www.us-proxy.org/"

// 	// Locator key is source url for List method
// 	l := Locator{
// 		Bucket: "finance-lake",
// 		Key:    filepath.Join("ingest/proxy/", strings.ReplaceAll(source, "/", "%2F"), ""),
// 	}

// 	locators, err := s.List(l)
// 	assert.Nil(t, err, "store List() has error")

// 	for _, l := range locators {
// 		file, _ := s.Read(l)
// 		// assert.Nil(t, err, "error in Read for key: "+file.Key)
// 		fmt.Println(file)
// 	}
// }
