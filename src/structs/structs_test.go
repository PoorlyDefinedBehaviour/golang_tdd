package structs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RectanglePerimeter(t *testing.T) {
	t.Parallel()

	rect := Rectangle{Width: 10.0, Height: 10.0}

	assert.Equal(t, 40.0, rect.Perimeter())
}

func Test_RectangleArea(t *testing.T) {
	t.Parallel()

	rect := Rectangle{Width: 12.0, Height: 6.0}

	assert.Equal(t, 72.0, rect.Area())
}

func Test_CircleArea(t *testing.T) {
	t.Parallel()

	circle := Circle{Radius: 10.0}

	assert.Equal(t, 314.1592653589793, circle.Area())
}
