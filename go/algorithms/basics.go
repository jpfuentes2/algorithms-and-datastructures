package algorithms

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// Shuffle does a Fisher-Yates shuffle
//
// Time O(n)
func Shuffle(slice []int) {
	for i := len(slice) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// Reverse a slice of int
//
// Time O(n)
func Reverse(s []int) {
	for i := len(s)/2 - 1; i >= 0; i-- {
		opp := len(s) - 1 - i
		s[i], s[opp] = s[opp], s[i]
	}
}

// BinarySearch does a binary search. It assumes the slice is
// already sorted.
//
// Time O(log n)
func BinarySearch(slice []int, search int) int {
	min, max := 0, len(slice)-1
	for min <= max {
		middle := (min + max) / 2

		if slice[middle] == search {
			return middle
		}

		if slice[middle] < search {
			min = middle + 1
		} else {
			max = middle - 1
		}
	}

	return -1
}
