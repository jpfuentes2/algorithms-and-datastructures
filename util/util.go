package datastructures

// Max gives the max of a vs b
func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// Min gives the min of a vs b
func Min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
