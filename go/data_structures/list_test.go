package data_structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewList(t *testing.T) {
	t.Parallel()

	list := SinglyLinkedList{}
	assert.True(t, list.IsEmpty())

	list.Insert(1)
	assert.False(t, list.IsEmpty())
}

func TestList(t *testing.T) {
	t.Parallel()

	abcs := []string{"a", "b", "c"}
	list := SinglyLinkedList{}

	for _, letter := range abcs {
		list.Insert(letter)
		assert.Equal(t, letter, list.Head().value)
	}

	expected := make([]string, 0, 3)
	list.Reverse()
	list.Each(func(n *Node) {
		expected = append(expected, n.value.(string))
	})

	assert.Equal(t, abcs, expected)
	assert.Equal(t, "a", list.Head().value)
	assert.Equal(t, "c", list.Tail().value)

	list.Delete()
	assert.Equal(t, "b", list.Head().value)

	list.Delete()
	list.Delete()
	assert.True(t, list.IsEmpty())
}
