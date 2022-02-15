package config

import (
	"fmt"
	"testing"

	"github.com/trietmnj/scraperCookie/internal/types"
)

func TestNewConfig(t *testing.T) {
	c, err := NewConfig(types.Json, "/workspaces/scraperCookie/test/config.json")
	fmt.Println(c)
	fmt.Println(err)
}
