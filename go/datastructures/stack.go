package datastructures

import (
	"errors"
	"sync"
)

// Stack is a thread-safe version of the stack abstract data type
// using a linked list
type Stack struct {
	list   *SinglyLinkedList
	length int

	mu *sync.RWMutex
}

// NewStack makes a Stack
func NewStack() *Stack {
	return &Stack{list: &SinglyLinkedList{}, mu: &sync.RWMutex{}}
}

// Push pushes an item onto the top of the stack increasing length by 1
// Time O(1)
func (s *Stack) Push(v interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.length++
	s.list.Insert(v)
}

// Pop removes and returns the top item of the stack reducing length by 1
// Time O(1)
func (s *Stack) Pop() (interface{}, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.length == 0 {
		return nil, errors.New("Pop")
	}

	s.length--
	return s.list.Delete().value, nil
}

// Peek returns the top item of the stack
// Time O(1)
func (s *Stack) Peek() (interface{}, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.length == 0 {
		return nil, errors.New("Pop")
	}

	return s.list.Head().value, nil
}

// IsEmpty ...
// Time O(1)
func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}

// Len gives the length of stack
// Time O(1)
func (s *Stack) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.length
}
