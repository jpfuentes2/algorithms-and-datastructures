package data_structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	t.Parallel()

	abcs := []string{"a", "b", "c"}
	stack := NewStack()
	assert.True(t, stack.IsEmpty())

	for _, letter := range abcs {
		stack.Push(letter)
		peek, err := stack.Peek()

		assert.NoError(t, err)
		assert.Equal(t, letter, peek)
	}

	assert.Equal(t, 3, stack.Len())
	assert.False(t, stack.IsEmpty())

	for _ = range abcs {
		_, err := stack.Pop()
		assert.NoError(t, err)
	}

	assert.True(t, stack.IsEmpty())
	assert.Equal(t, 0, stack.Len())
}
