package store

import (
	"fmt"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
)

var s = S3Store{}

func Test_upload(t *testing.T) {
	s.Init()
	err := s.StoreJson(
		aws.String("finance-lake"),
		aws.String("bronze/ingest/vic/ideasum-json/test.json"),
		strings.NewReader(
			`{"title":"Survey Test","description":"This is a description of the test survey","active":true}`,
		))
	fmt.Println(err)
}

func Test_keyExists(t *testing.T) {
	s.Init()
	exists, _ := s.KeyExists("finance-lake", "bronze/ingest/vic/ideasum-json/test.json")
	fmt.Println(exists)
	exists, _ = s.KeyExists("finance-lake", "bronze/ingest/vic/ideasum-json/test")
	fmt.Println(exists)
}
