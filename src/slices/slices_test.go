package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Sum(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 0, Sum([]int{}))

	assert.Equal(t, 1, Sum([]int{1}))

	assert.Equal(t, 0, Sum([]int{1, -1}))

	assert.Equal(t, 1, Sum([]int{1, -1, 1}))

	assert.Equal(t, 15, Sum([]int{1, 2, 3, 4, 5}))
}

func Test_SumAll(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 0, SumAll([]int{}, []int{}, []int{}))

	assert.Equal(t, 3, SumAll([]int{1}, []int{1}, []int{1}))

	assert.Equal(t, 2, SumAll([]int{1}, []int{-1}, []int{1}))
}
