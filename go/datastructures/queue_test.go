package datastructures

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

	abc := []string{"a", "b", "c"}
	queue = NewQueue()

	for i, letter := range abc {
		queue.Enqueue(letter)
		peek, err := queue.Peek()

		assert.NoError(t, err)
		assert.Equal(t, abc[i], peek)
	}

	assert.Equal(t, 3, queue.Len())
	assert.False(t, queue.IsEmpty())

	for i := range abc {
		actual, err := queue.Dequeue()

		assert.NoError(t, err)
		assert.Equal(t, abc[i], actual)
	}

	assert.True(t, queue.IsEmpty())
	assert.Equal(t, 0, queue.Len())
}

func TestErrEmptyQueue(t *testing.T) {
	t.Parallel()

	queue := NewQueue()
	_, err := queue.Dequeue()

	assert.Error(t, err)
	assert.Equal(t, ErrEmptyQueue, err)

	_, err = queue.Peek()

	assert.Error(t, err)
	assert.Equal(t, ErrEmptyQueue, err)
}
