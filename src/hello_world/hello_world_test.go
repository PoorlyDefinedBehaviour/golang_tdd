package helloworld

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	t.Parallel()

	tests := []struct {
		message  string
		name     string
		expected string
	}{
		{
			message:  "says hello to the name passed in ",
			name:     "John",
			expected: "Hello, John",
		},
		{
			message:  "returns the default message if name is empty",
			name:     "",
			expected: "Hello, worl1d",
		},
	}

	for _, tt := range tests {
		t.Log(tt.message)

		actual := Hello(tt.name)

		assert.Equal(t, tt.expected, actual)
	}
}
