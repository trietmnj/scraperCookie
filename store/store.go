package store

import (
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type IStore interface {
	Init()
	StoreJson(bucket *string, key *string, data io.ReadSeeker) error
	Read()
	KeyExists(k interface{}) (bool, error)
}

// https://john.dev/posts/2019-03-31-lambda-to-s3-golang.html
type S3Store struct {
	Uploader  *s3manager.Uploader
	s3Service *s3.S3
	region    *string
	creds     *credentials.Credentials
}

// type AWSConfig struct {
// 	AWS_REGION            string `envconfig:"AWS_REGION"`
// 	AWS_ACCESS_KEY_ID     string `envconfig:"AWS_ACCESS_KEY_ID"`
// 	AWS_SECRET_ACCESS_KEY string `envconfig:"AWS_SECRET_ACCESS_KEY"`
// }

//GetEnvWithKey : get env value
func GetEnvWithKey(key string) string {
	return os.Getenv(key)
}

func (s *S3Store) Init() {
	s.creds = credentials.NewEnvCredentials()
	s.region = aws.String("us-east-1")

	sess, err := session.NewSession(
		&aws.Config{
			Region:      s.region,
			Credentials: s.creds,
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	sess = session.Must(sess, err)
	s.Uploader = s3manager.NewUploader(sess)
	s.s3Service = s3.New(sess)
}

// TODO
func read(p []byte) (i int, err error) {
	return 0, nil
}

// key - location/fileName.json
func (s S3Store) StoreJson(
	bucket *string,
	key *string,
	data io.ReadSeeker,
) error {
	// uploader
	params := &s3manager.UploadInput{
		Bucket: bucket,
		Key:    key,
		Body:   data,
	}
	_, err := s.Uploader.Upload(params)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return err
}

// Check if key exists in bucket
func (s S3Store) KeyExists(bucket string, key string) (bool, error) {
	_, err := s.s3Service.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
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
