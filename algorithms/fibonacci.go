package algorithms

// Fibonacci sequence
//
//        | 0                if n = 0
// F(n) = | 1                if n = 1
//        | F(n-1) + F(n-2)  if n > 1

// FibonacciRecursive returns the
func FibonacciRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
}

func Fibonacci(n int64) int64 {
	var a, b int64
	a, b = 0, 1
	for i := int64(0); i < n; i++ {
		a, b = b, (a + b)
	}
	return a
}
