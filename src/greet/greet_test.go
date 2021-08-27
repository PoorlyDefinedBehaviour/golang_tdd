package greet

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreet(t *testing.T) {
	t.Parallel()

	buffer := bytes.Buffer{}

	Greet(&buffer, "John")

	expected := "Hello, John"

	assert.Equal(t, expected, buffer.String())
}
