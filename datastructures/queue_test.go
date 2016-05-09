package datastructures

import (
	"sort"
	"sync"
	"testing"

	"github.com/jpfuentes2/algorithms-and-datastructures/util"
	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
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
	queue := NewQueue()
	_, err := queue.Dequeue()

	assert.Error(t, err)
	assert.Equal(t, ErrEmptyQueue, err)

	_, err = queue.Peek()

	assert.Error(t, err)
	assert.Equal(t, ErrEmptyQueue, err)
}

func TestQueueThreadSafety(t *testing.T) {
	queue := NewQueue()
	ints := util.Range(0, 50)
	wg := sync.WaitGroup{}
	wg.Add(len(ints))

	for _, i := range ints {
		go func(x int) {
			queue.Enqueue(x)
			wg.Done()
		}(i)
	}

	wg.Wait()

	assert.False(t, queue.IsEmpty())
	assert.Equal(t, len(ints), queue.Len())

	wg = sync.WaitGroup{}
	wg.Add(len(ints))
	actual := struct {
		ints []int
		*sync.Mutex
	}{make([]int, 0, len(ints)), &sync.Mutex{}}

	for i := 0; i <= len(ints); i++ {
		go func() {
			if e, err := queue.Dequeue(); err == nil {
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

	assert.True(t, queue.IsEmpty())
	assert.Equal(t, 0, queue.Len())
	assert.Equal(t, ints, actual.ints)
}
