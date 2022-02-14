package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/trietmnj/scraperCookie/internal/types"
)

type AppConfig struct {
	ConfigType  types.ConfigSource
	JsonPath    string // path to config json
	StoreConfig        // store specific
}

type StoreConfig struct {
	S3Store      S3StoreConfig    `json:"s3store"`
	LocalStore   LocalStoreConfig `json:"localstore"`
	Repo         string           `json:"repo"`
	Proxy        bool             `json:"proxy"`
	StoreVariant store.Variant    `json:"storevariant"`
}

type S3StoreConfig struct {
	Bucket string `json:"bucket"`
}

type LocalStoreConfig struct {
	Path string `json:"path"`
}

// path - path to json config file
func NewConfig(v Variant, path string) (AppConfig, error) {
	var sc StoreConfig
	var err error
	switch v {
	case Json:
		jsonFile, err := os.Open(path)
		if err != nil {
			return AppConfig{}, err
		}
		defer jsonFile.Close()

		jsonByte, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			return AppConfig{}, err
		}

		var sc StoreConfig
		json.Unmarshal(jsonByte, &sc)
	default:
		return AppConfig{}, errors.New("invalid config variant")
	}
	return AppConfig{
		configVariant: v,
		JsonPath:      path,
		StoreConfig:   sc,
	}, err
}
