package config

import (
	"fmt"
	"testing"
)

func Test_NewConfig(t *testing.T) {
	c, err := NewConfig(Json, "./test/config.json")
	fmt.Println(c)
	fmt.Println(err)
}
