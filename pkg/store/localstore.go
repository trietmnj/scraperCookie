package store

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/trietmnj/scraperCookie/internal/util"
)

type LocalStore struct {
	StorePath string `envconfig:"storepath" required:"true"` // field name has to be exportable to work with envconfig
	file      *os.File
}

func (s *LocalStore) Read(l Locator) ([]byte, error) {
	filePath := filepath.Join(s.StorePath, l.Bucket, l.Key)
	return os.ReadFile(filePath)
}

func (s *LocalStore) Store(l Locator, data io.Reader) error {
	filePath := filepath.Join(s.StorePath, l.Bucket, l.Key)
	err := os.MkdirAll(util.Path(filePath), os.ModePerm)
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

func (s *LocalStore) KeyExists(l Locator) (bool, error) {
	filePath := filepath.Join(s.StorePath, l.Bucket, l.Key)
	if _, err := os.Stat(filePath); err == nil {
		return true, nil
	} else if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, fmt.Errorf("local store: unable to detect if key exists")
}

// List returns a list of files at the locator. Input Locator key has to be a folder.
func (s *LocalStore) List(l Locator) ([]Locator, error) {
	// TODO validate input locator is a folder
	var files []string
	root := filepath.Join(s.StorePath, l.Bucket, l.Key)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		} else if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	var ls []Locator
	for _, file := range files {
		ls = append(ls, Locator{
			Bucket: l.Bucket,
			Key:    strings.ReplaceAll(strings.ReplaceAll(file, s.StorePath+"/", ""), l.Bucket+"/", ""),
		})
	}
	return ls, err
}
