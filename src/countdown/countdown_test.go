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

func TestCount(t *testing.T) {
	t.Parallel()

	buffer := bytes.Buffer{}

	sleeperSpy := SleeperSpy{}

	Count(&buffer, &sleeperSpy)

	expected := `3
2
1
Go!`

	assert.Equal(t, expected, buffer.String())
	assert.Equal(t, 4, sleeperSpy.Calls)
}
