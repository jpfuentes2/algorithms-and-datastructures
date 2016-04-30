package data_structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	t.Parallel()

	queue := NewQueue()
	assert.True(t, queue.IsEmpty())
	assert.Equal(t, 0, queue.Len())

	queue.Enqueue(1)
	assert.False(t, queue.IsEmpty())
	assert.Equal(t, 1, queue.Len())

	abcs := []string{"a", "b", "c"}
	queue = NewQueue()

	for _, letter := range abcs {
		queue.Enqueue(letter)
		// peek, err := stack.Peek()

		// assert.NoError(t, err)
		// assert.Equal(t, letter, peek)
	}

	// assert.Equal(t, 3, stack.Len())
	// assert.False(t, stack.IsEmpty())

	// for _ = range abcs {
	// 	_, err := stack.Pop()
	// 	assert.NoError(t, err)
	// }

	// assert.True(t, stack.IsEmpty())
	// assert.Equal(t, 0, stack.Len())
}
