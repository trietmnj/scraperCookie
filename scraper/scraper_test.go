package downloader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendRequest(t *testing.T) {
	c := RequestConfig{
		Endpoint:  "/get",
		Type:      GET,
		URLParams: map[string]string{},
	}

	d := Downloader{
		ApiHost: "https://httpbin.org",
	}

	_, err := d.SendRequest(c)

	if assert.NoError(t, err) {
		assert.Equal(t, err, false)
	}
}
