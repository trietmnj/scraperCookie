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

// S3Store Init() using env credentials
type S3Store struct {
	uploader  *s3manager.Uploader
	s3Service *s3.S3
}

func (s *S3Store) Init() {
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
func (s *S3Store) Read(l Locator) []byte {
	fmt.Println(l)
	return []byte{}
}

// Locator args: bucket, key - location/fileName.json
func (s *S3Store) Store(
	l Locator,
	data io.Reader,
) error {
	key := l.Key
	if !strings.HasSuffix(key, ".json") {
		key += ".json"
	}
	params := &s3manager.UploadInput{
		Bucket: aws.String(l.Bucket),
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
func (s *S3Store) KeyExists(l Locator) (bool, error) {
	key := l.Key
	if !strings.HasSuffix(key, ".json") {
		key += ".json"
	}
	_, err := s.s3Service.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(l.Bucket),
		Key:    aws.String(key),
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
