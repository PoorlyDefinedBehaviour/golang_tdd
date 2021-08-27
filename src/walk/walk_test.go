package walk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStruct(t *testing.T) {
	t.Parallel()

	x := struct {
		name string
	}{
		"john",
	}

	expected := []string{"john"}

	var actual []string

	Struct(x, func(fieldName string) {
		actual = append(actual, fieldName)
	})

	assert.Equal(t, expected, actual)
}
