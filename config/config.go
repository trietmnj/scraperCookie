package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// type IConfigProvider interface {
// 	ProvideConfig() config
// }

type config struct {
	Bucket, DataSource, RepoName string
}

// type ApplicationConfig struct {
// 	AwsRegion string `json:"awsRegion"`
// 	AwsRegion string `json:"awsRegion"`
// }

// Load config from env vars
func NewConfig() config {
	var c config
	if err := envconfig.Process("", &c); err != nil {
		log.Fatal(err.Error())
	}
	return c
}

// func GetConfig() (Config, error) {
// 	sess, err := session.NewSession()
// 	svc := session.Must(sess, err)
// 	ssmsvc := &SSM{ssm.New(sess)}
// }

// func (c Config) String() {
// 	return string(c)
// }
