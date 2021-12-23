package store

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// Locator or index to find data in store
// l[0] = bucket
// l[1] = source
// l[2] = repo
// l[3] = url
type Locator []string

// Base interface, should not be fed directly to scraper
type IStore interface {
	Init()
	Store(l Locator, data io.Reader) error
	Read(l Locator) []byte
}

type IS3JsonStore interface {
	IStore
	KeyExists(l Locator) (bool, error) // check if data exists without reading the entire data
}

type S3JsonStore struct {
	uploader  *s3manager.Uploader
	s3Service *s3.S3
}

func (s *S3JsonStore) Init() {
	creds := credentials.NewEnvCredentials()
	region := aws.String("us-east-1")

	sess, err := session.NewSession(
		&aws.Config{
			Region:      region,
			Credentials: creds,
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	sess = session.Must(sess, err)
	s.uploader = s3manager.NewUploader(sess)
	s.s3Service = s3.New(sess)
}

// TODO
func (s *S3JsonStore) Read(l Locator) []byte {
	fmt.Println(l)
	return []byte{}
}

// Locator args: bucket, key - location/fileName.json
func (s *S3JsonStore) Store(
	l Locator,
	data io.Reader,
) error {
	var key string
	key = "bronze/ingest/" + l[1] + "/" + l[2] + "/" + l[3]
	if !strings.HasSuffix(key, ".json") {
		key += ".json"
	}
	params := &s3manager.UploadInput{
		Bucket: aws.String(l[0]),
		Key:    aws.String(key),
		Body:   data,
	}
	_, err := s.uploader.Upload(params)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return err
}

// Check if key exists in bucket
func (s *S3JsonStore) KeyExists(l Locator) (bool, error) {
	_, err := s.s3Service.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(l[0]),
		Key:    aws.String(l[1]),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case "NotFound": // s3.ErrCodeNoSuchKey does not work, aws is missing this error code so we hardwire a string
				return false, nil
			default:
				return false, err
			}
		}
		return false, err
	}
	return true, nil
}
