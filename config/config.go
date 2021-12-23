package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type IConfigProvider interface {
	ProvideConfig() config
}

type config struct {
	Bucket, DataSource, RepoName string
}

// Load config from env vars
func Init() config {
	var c config
	if err := envconfig.Process("", &c); err != nil {
		log.Fatal(err.Error())
	}
	return c
}

// func (c Config) String() {
// 	return string(c)
// }
