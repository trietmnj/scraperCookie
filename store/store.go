package store

import ()

type Config struct {
	AwsS3Region string
	AwsS3Bucket string
}

// func Init()
// 	var cfg Config
// 	if err := envconfig.Process("", &cfg); err != nil {
// 		log.Fatal(err.Error())
// 	}
