package helloworld

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	t.Parallel()

	t.Run("says hello to the name passed in", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, "Hello, John", Hello("John", Languages.English))
	})

	t.Run("returns a default message if name is empty", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, "Hello, world", Hello("", Languages.English))
	})

	t.Run("in spanish", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, "Hola, John", Hello("John", Languages.Spanish))
	})
}
