package sorting

/*
   Insertion sort

   http://www.sorting-algorithms.com/insertion-sort
   https://en.wikipedia.org/wiki/Insertion_sort
   http://visualgo.net/sorting

   - Stable
   - In-place
   - Space O(1)
   - Time (worst case): O(n^2)
   - Adaptive: O(n) time when nearly sorted
*/

// InsertionSort sorts the slice using insertion sort algorithm
func InsertionSort(ints []int) {
	for i := 1; i < len(ints); i++ {
		x := ints[i]
		j := i - 1
		for ; j >= 0 && ints[j] > x; j-- {
			ints[j+1] = ints[j]
		}
		ints[j+1] = x
	}
}
