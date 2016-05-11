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
// Space O(1)
func Shuffle(slice []int) {
	// base case: cannot shuffle empty or 1-element slice
	if len(slice) <= 1 {
		return
	}

	for i := len(slice) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// Reverse a slice of int
//
// Time O(n)
// Space O(1)
func Reverse(s []int) {
	// base case: cannot reverse empty or 1-element slice
	if len(s) <= 1 {
		return
	}

	for i := len(s)/2 - 1; i >= 0; i-- {
		opp := len(s) - 1 - i
		s[i], s[opp] = s[opp], s[i]
	}
}

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
		}

		if slice[middle] < search {
			min = middle + 1
		} else {
			max = middle - 1
		}
	}

	return -1
}

// IsPalindrome returns true if the given string is a palindrome ignoring whitespace,
// punctuation, & case sensitivity
// Time O(n)
// Space O(1)
func IsPalindrome(str string) bool {
	n := len(str)

	// base case: empty or single char string are palindromes
	if n == 0 || n == 1 {
		return true
	}

	for i := 0; i < n/2; i++ {
		front := str[i]
		back := str[n-1]

		if front != back {
			return false
		}

		n--
	}

	return true
}

// IsPalindromeRecursive does the same as IsPalindrome using recursion
// Time O(n)
// Space O(1)
func IsPalindromeRecursive(str string) bool {
	var recurse func(s string) bool

	recurse = func(s string) bool {
		n := len(s)
		// base case: empty or single char string are palindromes
		if n == 0 || n == 1 {
			return true
		} else if s[0] != s[n-1] {
			return false
		} else {
			return recurse(s[1 : n-1])
		}
	}

	return recurse(str)
}
