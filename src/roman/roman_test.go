package roman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromArabic(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "I", FromArabic(1))
}
