package store

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/trietmnj/scraperCookie/config"
)

// https://john.dev/posts/2019-03-31-lambda-to-s3-golang.html
type Store struct {
	Uploader *s3manager.Uploader
	Config   config.Config // configs specific to s3 access
}

func (s Store) Init() {
	s.Config.Init()
	credentials.NewEnvCredentials()
	sess := session.Must(session.NewSession())
	s.Uploader = s3manager.NewUploader(sess)
}

func (s Store) Upload(data io.Reader) {
	s.Uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s.Config.AwsS3Bucket),
		Key:    aws.String("test.json"), // file name
		Body:   data,
	})
}
