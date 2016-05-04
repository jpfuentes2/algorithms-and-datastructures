package datastructures

import (
	dblList "container/list"
	"errors"
	"sync"
)

// EmptyQueueError is the error when a queue is unexpectedly empty
var ErrEmptyQueue = errors.New("Queue is empty!")

// Queue is a thread-safe FIFO queue using a doubly linked list for O(1) operations
type Queue struct {
	list *dblList.List
	mu   *sync.RWMutex
}

// NewQueue creates an empty queue
func NewQueue() *Queue {
	return &Queue{list: dblList.New(), mu: &sync.RWMutex{}}
}

// Enqueue adds an element to the back of the queue.
// Time O(1)
// Space O(1)
func (q *Queue) Enqueue(v interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.list.PushBack(v)
}

// Dequeue removes and returns front element. EmptyQueueError if empty.
// Time O(1)
// Space O(1)
func (q *Queue) Dequeue() (interface{}, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.list.Len() == 0 {
		return nil, ErrEmptyQueue
	}

	return q.list.Remove(q.list.Front()), nil
}

// Peek gives the front element without modifying the queue.
// EmptyQueueError if empty.
// Time O(1)
// Space O(1)
func (q *Queue) Peek() (interface{}, error) {
	q.mu.RLock()
	defer q.mu.RUnlock()

	if q.list.Len() == 0 {
		return nil, ErrEmptyQueue
	}

	return q.list.Back().Value, nil
}

// IsEmpty is true if there's nothing in the queue
// Time O(1)
// Space O(1)
func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}

// Len gives the length of this queue
// Time O(1)
// Space O(1)
func (q *Queue) Len() int {
	q.mu.RLock()
	defer q.mu.RUnlock()

	return q.list.Len()
}
