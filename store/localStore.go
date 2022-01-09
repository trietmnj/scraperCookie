package store

import (
	"bytes"
	"io"
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
)

type localStore struct {
	storePath string `envconfig:""`
	file      *os.File
}

// Init() generates StorePath from env var LOCAL_STOREPATH
func (s *localStore) Init() {
	var c localStore
	// TODO something wrong with processing env var to configs
	err := envconfig.Process("local", &c)
	if err != nil {
		log.Fatal(err.Error())
	}
}

// Read is unimplemented
func (s *localStore) Read(l Locator) []byte {
	return []byte{}
}

func (s *localStore) Store(l Locator, data io.Reader) error {
	file, err := os.Create(s.storePath + l.Key)
	if err != nil {
		return err
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(data)
	file.Write(buf.Bytes())

	return err
}
