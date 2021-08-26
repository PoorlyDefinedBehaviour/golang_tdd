package dictionary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Search(t *testing.T) {
	t.Parallel()

	dict := map[string]string{
		"a": "b",
	}

	expected := "b"

	actual := Search(dict, "a")

	assert.Equal(t, expected, actual)
}
