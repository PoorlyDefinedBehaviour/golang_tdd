package countdown

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SleeperSpy struct {
	Calls int
}

func (sleeper *SleeperSpy) Sleep() {
	sleeper.Calls++
}

type CountdownOperationsSpy struct {
	Calls []string
}

const (
	write = "write"
	sleep = "sleep"
)

func (spy *CountdownOperationsSpy) Sleep() {
	spy.Calls = append(spy.Calls, sleep)
}

func (spy *CountdownOperationsSpy) Write(p []byte) (int, error) {
	spy.Calls = append(spy.Calls, write)
	return 0, nil
}

func TestCount(t *testing.T) {
	t.Parallel()

	t.Run("writes expected values", func(t *testing.T) {
		t.Parallel()

		buffer := bytes.Buffer{}

		sleeperSpy := SleeperSpy{}

		Count(&buffer, &sleeperSpy)

		expected := "3\n2\n1\nGo!"

		assert.Equal(t, expected, buffer.String())
		assert.Equal(t, 4, sleeperSpy.Calls)
	})

	t.Run("sleeps before every print", func(t *testing.T) {
		t.Parallel()

		spy := CountdownOperationsSpy{}

		expected := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		Count(&spy, &spy)

		assert.Equal(t, expected, spy.Calls)
	})
}
