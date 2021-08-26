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

func Test_Contains(t *testing.T) {
	t.Parallel()

	tests := []struct {
		dict     T
		key      string
		expected bool
	}{
		{
			dict:     T{},
			key:      "a",
			expected: false,
		},
		{
			dict: T{
				"a": "b",
			},
			key:      "c",
			expected: false,
		},
		{
			dict: T{
				"a": "b",
			},
			key:      "a",
			expected: true,
		},
	}

	for _, tt := range tests {
		actual := tt.dict.Contains(tt.key)

		assert.Equal(t, tt.expected, actual)
	}

}
