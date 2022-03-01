package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trietmnj/scraperCookie/internal/types"
)

func TestNewConfigLocal(t *testing.T) {
	c, err := NewConfig(types.JsonConfigSource, "/workspaces/scraperCookie/test/config-local.json")
	fmt.Println(c)
	assert.Nil(t, err)
}

func TestNewConfigS3(t *testing.T) {
	c, err := NewConfig(types.JsonConfigSource, "/workspaces/scraperCookie/test/config-s3.json")
	fmt.Println(c)
	assert.Nil(t, err)
}
