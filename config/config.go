package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// type IConfigProvider interface {
// 	ProvideConfig() config
// }

// type config struct {
// 	Bucket, DataSource, RepoName string
// }

type Config struct {
	Bucket    string `json:"bucket"`
	Repo      string `json:"repo"`
	LocalPath string `json"localPath"`
}

// type ApplicationConfig struct {
// 	AwsRegion string `json:"awsRegion"`
// 	AwsRegion string `json:"awsRegion"`
// }

// Load config from env vars
// func NewConfig() config {
// 	var c config
// 	if err := envconfig.Process("", &c); err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	return c
// }

// Load config from config.json
// TODO refactor config to select between env var and json
func NewConfig(path string) (Config, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	defer jsonFile.Close()

	jsonByte, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return Config{}, err
	}

	var c Config
	json.Unmarshal(jsonByte, &c)
	return c, err
}
