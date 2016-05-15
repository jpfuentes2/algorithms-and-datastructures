package sorting

/*
   Selection sort

   http://www.sorting-algorithms.com/selection-sort
   https://en.wikipedia.org/wiki/Selection_sort
   http://visualgo.net/sorting

   - Not Stable
   - In-place
   - Space O(1)
   - Time (all cases): O(n^2)
   - Not adaptive
*/

// SelectionSort sorts the slice using selection sort algorithm
func SelectionSort(ints []int) {
	n := len(ints)
	for i := 0; i < n-1; i++ {
		currMin := i

		for j := i + 1; j < n; j++ {
			if ints[j] < ints[currMin] {
				currMin = j
			}
		}

		if currMin != i {
			ints[currMin], ints[i] = ints[i], ints[currMin]
		}
	}
}
