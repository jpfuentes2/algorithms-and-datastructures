package datastructures

import (
	"sort"
	"sync"
	"testing"

	"github.com/jpfuentes2/algorithms-and-datastructures/util"
	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
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

func TestErrEmptyStack(t *testing.T) {
	stack := NewStack()
	_, err := stack.Pop()

	assert.Error(t, err)
	assert.Equal(t, ErrEmptyStack, err)

	_, err = stack.Peek()

	assert.Error(t, err)
	assert.Equal(t, ErrEmptyStack, err)
}

func TestStackThreadSafety(t *testing.T) {
	stack := NewStack()
	ints := util.Range(0, 50)
	wg := sync.WaitGroup{}
	wg.Add(len(ints))

	for _, i := range ints {
		go func(x int) {
			stack.Push(x)
			wg.Done()
		}(i)
	}

	wg.Wait()

	assert.False(t, stack.IsEmpty())
	assert.Equal(t, len(ints), stack.Len())

	wg = sync.WaitGroup{}
	wg.Add(len(ints))
	actual := struct {
		ints []int
		*sync.Mutex
	}{make([]int, 0, len(ints)), &sync.Mutex{}}

	for i := 0; i <= len(ints); i++ {
		go func() {
			if e, err := stack.Pop(); err == nil {
				actual.Lock()
				actual.ints = append(actual.ints, (e).(int))
				actual.Unlock()
				wg.Done()
			}
		}()
	}

	wg.Wait()

	sort.Ints(ints)
	sort.Ints(actual.ints)

	assert.True(t, stack.IsEmpty())
	assert.Equal(t, 0, stack.Len())
	assert.Equal(t, ints, actual.ints)
}
