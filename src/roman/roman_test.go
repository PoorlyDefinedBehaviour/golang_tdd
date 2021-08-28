package roman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromArabic(t *testing.T) {
	t.Parallel()

	tests := []struct {
		n        int
		expected string
	}{
		{
			n:        1,
			expected: "I",
		},
		{
			n:        2,
			expected: "II",
		},
		{
			n:        3,
			expected: "III",
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, FromArabic(tt.n))
	}
}
