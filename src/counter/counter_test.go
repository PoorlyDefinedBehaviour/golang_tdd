package counter

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncrement(t *testing.T) {
	t.Parallel()

	t.Run("counter starts at the value provided to New", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			counter  *T
			expected int
		}{
			{
				counter:  New(-1),
				expected: -1,
			},
			{
				counter:  New(0),
				expected: 0,
			},
			{
				counter:  New(1),
				expected: 1,
			},
			{
				counter:  New(3),
				expected: 3,
			},
		}

		for _, tt := range tests {
			assert.Equal(t, tt.expected, tt.counter.Value())
		}
	})

	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		t.Parallel()

		counter := New(0)

		counter.Increment()
		counter.Increment()
		counter.Increment()

		assert.Equal(t, 3, counter.Value())
	})

	t.Run("runs safely concurrently", func(t *testing.T) {
		t.Parallel()

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
