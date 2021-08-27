package counter

import (
	"sync"
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

	t.Run("runs safely concurrently", func(t *testing.T) {
		expected := 5000

		counter := New(0)

		waitGroup := sync.WaitGroup{}
		waitGroup.Add(expected)

		for i := 0; i < expected; i++ {
			go func() {
				counter.Increment()
				waitGroup.Done()
			}()
		}

		waitGroup.Wait()

		assert.Equal(t, expected, counter.Value())
	})
}
