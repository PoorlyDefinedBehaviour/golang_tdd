package iteration

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Repeat(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "", Repeat("a", 0))
	assert.Equal(t, "a", Repeat("a", 1))
	assert.Equal(t, "aaaaa", Repeat("a", 5))
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	fmt.Println(Repeat("a", 3))
	// Output: aaa
}
