package concurrency

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckWebsites(t *testing.T) {
	t.Parallel()

	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	expected := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	actual := CheckWebsites(
		func(url string) bool {
			if url == "waat://furhurterwe.geds" {
				return false
			}
			return true
		},
		websites,
	)

	assert.Equal(t, expected, actual)
}
