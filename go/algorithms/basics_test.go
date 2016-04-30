package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	t.Parallel()

	nums := []int{1, 2, 3, 4, 5}
	Reverse(nums)
	assert.Equal(t, nums, []int{5, 4, 3, 2, 1})
}

// need to shuffle N times, bin results, and ensure
// uniform distribution via Pearson Chi-squared test
func TestShuffle(t *testing.T) {
	t.Parallel()

	n := 4
	iterations := 100000
	var bins = make(map[string]int)
	var nums = make([]int, 0, n)

	for i := 0; i <= n; i++ {
		nums = append(nums, i)
	}

	for i := 0; i <= iterations; i++ {
		tmp := make([]int, n)
		copy(tmp, nums)

		Shuffle(tmp)
		bins[fmt.Sprint(tmp)]++
	}

	assert.Fail(t, ":(")
}

func TestBinarySearch(t *testing.T) {
	t.Parallel()

	nums := []int{1, 2, 3, 4, 5}
	for i := range nums {
		assert.Equal(t, i, BinarySearch(nums, i+1))
	}

	assert.Equal(t, -1, BinarySearch(nums, 0))
	assert.Equal(t, -1, BinarySearch(nums, 10))
}
