package proxy

import (
	"fmt"
	"testing"

	"github.com/trietmnj/scraperCookie/store"
)

func TestProxy(t *testing.T) {
	s, _ := store.NewStore("local")
	l := store.Locator{
		Key:    "https://www.us-proxy.org/",
		Bucket: "finance-lake",
	}
	pp, err := NewProxyFunction(s, l)
	fmt.Println(err)
	fmt.Println(pp)
}
