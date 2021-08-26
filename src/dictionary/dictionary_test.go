package dictionary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Search(t *testing.T) {
	t.Parallel()

	dict := T{
		"a": "b",
	}

	expected := "b"

	actual := dict.Search("a")

	assert.Equal(t, expected, actual)
}
