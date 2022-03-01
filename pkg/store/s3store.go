package store

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/trietmnj/scraperCookie/pkg/config"
)

// s3Store Init() using env credentials
type s3Store struct {
	uploader  *s3manager.Uploader
	s3Service *s3.S3
}

// c should be of type config.S3StoreConfig
// aws credential should be in env vars: AWS_ACCESS_KEY_ID + AWS_SECRET_ACCESS_KEY
func (s *s3Store) init(c interface{}) error {
	coercedC, ok := c.(config.S3StoreConfig)
	if !ok {
		return errors.New("s3store init: unable to read s3 store config")
	}
	creds := credentials.NewEnvCredentials()

	sess, err := session.NewSession(
		&aws.Config{
			Region:      aws.String(coercedC.Region),
			Credentials: creds,
		},
	)
	if err != nil {
		return err
	}

	sess = session.Must(sess, err)
	s.uploader = s3manager.NewUploader(sess)
	s.s3Service = s3.New(sess)
	return err
}

// TODO
func (s *s3Store) Read(l iLocator) ([]byte, error) {
	return []byte{}, nil
}

func (s *s3Store) Store(l iLocator, data io.Reader) error {
	params := &s3manager.UploadInput{
		Bucket: aws.String(l.Container()),
		Key:    aws.String(l.File()),
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
func (s *s3Store) KeyExists(l iLocator) (bool, error) {
	_, err := s.s3Service.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(l.Container()),
		Key:    aws.String(l.File()),
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

// TODO
func (s *s3Store) List(l iLocator) ([]Locator, error) {
	return []Locator{}, nil
}
