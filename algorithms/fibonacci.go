package algorithms

// Fibonacci computes the nth Fibonacci number
//
//   	       | 0                if n = 0
//   	F(n) = | 1                if n = 1
//   	       | F(n-1) + F(n-2)  if n > 1
//
// Time O(N) | Space O(1)
func Fibonacci(n int64) int64 {
	var a, b int64
	a, b = 0, 1
	for i := int64(0); i < n; i++ {
		a, b = b, (a + b)
	}
	return a
}

// FibonacciRecursive computes the nth Fibonacci number
//
// Time O(2^n)
func FibonacciRecursive(n int64) int64 {
	if n <= 1 {
		return n
	}
	return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
}
