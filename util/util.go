package util

import "math"

// Max gives the max of a vs b
//
// Max(a,b) =  | a if a > b
//             | b otherwise
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min gives the min of a vs b
//
// Min(a,b) =  | a if a < b
//             | b otherwise
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Range creates a slice of integers inclusively [a,b]
func Range(start, end int) []int {
	n := int(math.Abs(float64(end - start)))
	if n < 0 {
		return []int{}
	}

	ints := make([]int, n+1)
	j := start
	for i := 0; i <= n; i++ {
		ints[i] = j
		j++
	}
	return ints
}
