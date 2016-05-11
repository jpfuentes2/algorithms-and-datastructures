package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}

	cases := []struct {
		Haystack []int
		Needle   int
		Expected int
	}{
		{nums, 1, 0},
		{nums, 3, 2},
		{nums, 5, 4},
		{nums, 0, -1},
		{nums, 6, -1},
		{[]int{}, 6, -1},
	}

	for _, tc := range cases {
		actual := BinarySearch(tc.Haystack, tc.Needle)
		actualRecursive := BinarySearchRecursive(tc.Haystack, tc.Needle)

		assert.Equal(t, tc.Expected, actual)
		assert.Equal(t, tc.Expected, actualRecursive)
	}
}
