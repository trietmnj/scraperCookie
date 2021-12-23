package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type IConfigProvider interface {
	ProvideConfig() Config
}

type Config struct {
}

// Load config from env vars
func (c Config) Init() {
	if err := envconfig.Process("", &c); err != nil {
		log.Fatal(err.Error())
	}
}

// func (c Config) String() {
// 	return string(c)
// }
