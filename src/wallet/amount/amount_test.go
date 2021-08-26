package amount

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	t.Parallel()

	t.Run("shouldn't be able to create a negative amount", func(t *testing.T) {
		t.Parallel()

		_, err := New(-1)

		assert.True(t, errors.Is(err, ErrInvalidAmount))
	})

	t.Run("should be able to create valid amount", func(t *testing.T) {
		t.Parallel()

		_, err := New(0)

		assert.Nil(t, err)

		_, err = New(1)

		assert.Nil(t, err)

		_, err = New(10)

		assert.Nil(t, err)
	})
}

func Test_Add(t *testing.T) {
	t.Parallel()

	tests := []struct {
		a        int
		b        int
		expected int
	}{
		{
			a:        10,
			b:        10,
			expected: 20,
		},
		{
			a:        0,
			b:        1,
			expected: 1,
		},
		{
			a:        0,
			b:        0,
			expected: 0,
		},
	}

	for _, tt := range tests {
		a, _ := New(tt.a)

		b, _ := New(tt.b)

		expected, _ := New(tt.expected)

		actual := Add(a, b)

		assert.Equal(t, expected, actual)
	}
}

func Test_Int(t *testing.T) {
	t.Parallel()

	amount, err := New(3)

	assert.Nil(t, err)

	assert.Equal(t, 3, Int(amount))
}

func TestT_String(t *testing.T) {
	t.Parallel()

	amount, err := New(5)

	assert.Nil(t, err)

	expected := "Amount(5)"

	assert.Equal(t, expected, amount.String())
}

func TestSub(t *testing.T) {
	t.Parallel()

	a, _ := New(10)

	b, _ := New(5)

	expected, _ := New(5)

	actual := Sub(a, b)

	assert.Equal(t, expected, actual)
}
