package structs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Perimeter(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 40.0, Perimeter(Rectangle{Width: 10.0, Height: 10.0}))
}

func Test_Area(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 72.0, Area(Rectangle{Width: 12.0, Height: 6.0}))
}
