package configs

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Bucket    string `json:"bucket"`
	Repo      string `json:"repo"`
	LocalPath string `json:"local-path"`
}

// Load config from config.json
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
