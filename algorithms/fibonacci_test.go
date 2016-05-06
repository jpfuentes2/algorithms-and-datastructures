package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	cases := []struct {
		N        int64
		Expected int64
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{10, 55},
	}

	for _, tc := range cases {
		assert.Equal(t, int(tc.Expected), FibonacciRecursive(int(tc.N)))
		assert.Equal(t, tc.Expected, Fibonacci(tc.N))
	}
}
