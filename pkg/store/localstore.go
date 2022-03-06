package store

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/trietmnj/scraperCookie/internal/types"
	"github.com/trietmnj/scraperCookie/pkg/config"
)

type LocalStore struct {
	BaseDirectory string
}

// c should be of type config.StoreConfig
func (s *LocalStore) init(c interface{}) error {
	coercedC, ok := c.(config.StoreConfig)
	s.BaseDirectory = coercedC.Path
	if !ok {
		return errors.New("localstore init: unable to read store config")
	}
	return nil
}

func (s *LocalStore) Read(l iLocator) ([]byte, error) {
	filePath := filepath.Join(l.Path(), l.File())
	return os.ReadFile(filePath)
}

func (s *LocalStore) Store(l iLocator, data io.Reader) error {
	path := filepath.Join(s.BaseDirectory, l.Path())
	filePath := filepath.Join(path, l.File())
	err := os.MkdirAll(path, os.ModePerm)
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

func (s *LocalStore) KeyExists(l iLocator) (bool, error) {
	filePath := filepath.Join(s.BaseDirectory, l.Path(), l.File())
	if _, err := os.Stat(filePath); err == nil {
		return true, nil
	} else if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, fmt.Errorf("local store: unable to detect if key exists")
}

// List returns a list of files at the locator. Input Locator key has to be a folder.
func (s *LocalStore) List(l iLocator) ([]Locator, error) {
	// TODO validate input locator is a folder
	var files []string
	root := filepath.Join(s.BaseDirectory, l.Path())
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
		dir, filename := filepath.Split(file)
		localLocator := LocalLocation{
			Path: dir,
			File: filename,
		}
		locator := Locator{
			StoreType: types.LocalStore,
			Local:     localLocator,
		}
		ls = append(ls, locator)
	}
	return ls, err
}
