package helloworld

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "Hello, John", Hello("John"))
}
