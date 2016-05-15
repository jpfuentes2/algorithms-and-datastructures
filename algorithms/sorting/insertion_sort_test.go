package sorting

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jpfuentes2/algorithms-and-datastructures/algorithms"
	"github.com/jpfuentes2/algorithms-and-datastructures/util"
)

// TestInsertionSort
func TestInsertionSort(t *testing.T) {
	cases := []struct {
		Ints []int
	}{
		{util.Range(0, 0)},
		{util.Range(0, 10)},
		{util.Range(0, 100)},
		{util.Range(5, 6)},
		{util.Range(-100, 100)},
		{[]int{5}},
		{[]int{5, 5}},
	}

	for _, tc := range cases {
		mine := make([]int, len(tc.Ints))
		theirs := make([]int, len(tc.Ints))
		original := make([]int, len(tc.Ints))

		copy(original, tc.Ints)
		algorithms.Shuffle(tc.Ints)
		copy(mine, tc.Ints)
		copy(theirs, tc.Ints)

		InsertionSort(mine)
		sort.Ints(theirs)

		assert.Equal(t, theirs, mine)
		assert.Equal(t, original, mine)
	}
}
