package datastructures

import (
	dblList "container/list"
	"sync"
)

type Queue struct {
	list *dblList.List
	mu   *sync.RWMutex
}

func NewQueue() *Queue {
	return &Queue{list: dblList.New(), mu: &sync.RWMutex{}}
}

func (q *Queue) Enqueue(v interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.list.PushBack(v)
}

func (q *Queue) Dequeue() (interface{}, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	return q.list.Remove(q.list.Front()), nil
}

func (q *Queue) Peek() (interface{}, error) {
	q.mu.RLock()
	defer q.mu.RUnlock()

	return q.list.Front(), nil
}

func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}

func (q *Queue) Len() int {
	q.mu.RLock()
	defer q.mu.RUnlock()

	return q.list.Len()
}
