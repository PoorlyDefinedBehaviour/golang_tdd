package integers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Add(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 4, Add(2, 2))
}

// This appears in the documentation if we use godoc
func ExampleAdd() {
	fmt.Println(Add(1, 5))
	// Output: 6
}
