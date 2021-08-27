package countdown

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCount(t *testing.T) {
	buffer := bytes.Buffer{}

	Count(&buffer)

	expected := `3
2
1
Go!`

	assert.Equal(t, expected, buffer.String())
}
