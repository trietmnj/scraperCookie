package config

import (
	"fmt"
	"testing"

	"github.com/trietmnj/scraperCookie/internal/types"
)

func TestNewConfigLocal(t *testing.T) {
	c, err := NewConfig(types.JsonConfigSource, "/workspaces/scraperCookie/test/config-local.json")
	fmt.Println(c)
	fmt.Println(err)
}

func TestNewConfigS3(t *testing.T) {
	c, err := NewConfig(types.JsonConfigSource, "/workspaces/scraperCookie/test/config-s3.json")
	fmt.Println(c)
	fmt.Println(err)
}
