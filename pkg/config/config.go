package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/trietmnj/scraperCookie/internal/types"
)

type AppConfig struct {
	ConfigSourceType types.ConfigSource
	JsonPath         string // path to config json
	StoreConfig
	ScraperConfig
}

type StoreConfig struct {
	S3Store    S3StoreConfig    `json:"s3store"`
	LocalStore LocalStoreConfig `json:"localstore"`
	Proxy      bool             `json:"proxy"`
	StoreType  types.Store      `json:"storetype"`
	Path       string           `json:"path" envconfig:"storepath" required:"true"`
}

type S3StoreConfig struct {
	Bucket string `json:"bucket"`
	Region string `json:"region"`
}

type LocalStoreConfig struct {
}

type ScraperConfig struct {
	DataType types.Data
}

// path - path to json config file
func NewConfig(v types.ConfigSource, path string) (AppConfig, error) {
	var sc StoreConfig
	var err error
	switch v {
	case types.JsonConfigSource:
		if path == "" {
			return AppConfig{}, errors.New("invalid json file path: " + path)
		}
		jsonFile, err := os.Open(path)
		if err != nil {
			return AppConfig{}, err
		}
		defer jsonFile.Close()
		jsonByte, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			return AppConfig{}, err
		}
		json.Unmarshal(jsonByte, &sc)

		// TODO add parsing config from env
	default:
		return AppConfig{}, errors.New("invalid config variant")
	}
	c := AppConfig{
		ConfigSourceType: v,
		JsonPath:         path,
		StoreConfig:      sc,
	}
	return c, err
}
