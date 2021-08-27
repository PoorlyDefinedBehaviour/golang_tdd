package counter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncrement(t *testing.T) {
	t.Parallel()

	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := New(0)

		counter.Increment()
		counter.Increment()
		counter.Increment()

		assert.Equal(t, 3, counter.Value())
	})
}
