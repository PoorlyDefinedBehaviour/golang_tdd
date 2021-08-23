package integers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Add(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 4, Add(2, 2))
}
