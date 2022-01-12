package store

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/kelseyhightower/envconfig"
	"github.com/trietmnj/scraperCookie/utils"
)

// TODO should local store really take config from env
// LOCAL_STOREPATH has to be available as an env var
type localStore struct {
	StorePath string `envconfig:"storepath" required:"true"` // field name has to be exportable to work with envconfig
	file      *os.File
}

// Init() generates StorePath from env var LOCAL_STOREPATH
func (s *localStore) Init() {
	var c localStore // temporary var to work with envconfig instead of direct mutation on reciever
	err := envconfig.Process("local", &c)
	if err != nil {
		log.Fatal(err.Error())
	}
	if c.StorePath == "" {
		log.Fatal("local store: unable to parse config from env")
	}
	s.StorePath = c.StorePath
}

// Read is unimplemented
func (s *localStore) Read(l Locator) []byte {
	return []byte{}
}

func (s *localStore) Store(l Locator, data io.Reader) error {
	filePath := filepath.Join(s.StorePath, l.Bucket, l.Key)
	err := os.MkdirAll(utils.Path(filePath), os.ModePerm)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(data)

	bytesData := buf.Bytes()
	_, err = file.Write(bytesData)

	return err
}

func (s *localStore) KeyExists(l Locator) (bool, error) {
	filePath := filepath.Join(s.StorePath, l.Bucket, l.Key)
	if _, err := os.Stat(filePath); err == nil {
		return true, nil
	} else if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, fmt.Errorf("local store: unable to detect if key exists")
}