package algorithms

// BinarySearch does a binary search. It assumes the slice is
// already sorted.
//
// Time O(log n)
// Space O(1)
func BinarySearch(slice []int, search int) int {
	min, max := 0, len(slice)-1
	for min <= max {
		middle := (min + max) / 2

		if slice[middle] == search {
			return middle
		} else if slice[middle] < search {
			min = middle + 1
		} else {
			max = middle - 1
		}
	}

	return -1
}

// BinarySearchRecursive does the same as BinarySearch except recursively.
//
// Time O(log n)
// Space O(1)
func BinarySearchRecursive(slice []int, search int) int {
	var recurse func(min, max int) int

	recurse = func(min, max int) int {
		if min > max {
			return -1
		}

		middle := (min + max) / 2

		if slice[middle] == search {
			return middle
		} else if slice[middle] < search {
			return recurse(middle+1, max)
		} else {
			return recurse(min, middle-1)
		}
	}

	return recurse(0, len(slice)-1)
}
