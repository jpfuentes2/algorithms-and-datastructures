package datastructures

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
