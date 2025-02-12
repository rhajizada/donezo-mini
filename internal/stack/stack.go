package stack

import (
	"sync"
)

type Stack[T any] struct {
	ch chan T
	mu sync.Mutex
}

func New[T any](capacity int) *Stack[T] {
	return &Stack[T]{ch: make(chan T, capacity)}
}

// Push adds an element to the stack
func (s *Stack[T]) Push(value T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.ch <- value
}

// Pop removes and returns an element from the stack
func (s *Stack[T]) Pop() (T, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	select {
	case value := <-s.ch:
		return value, true
	default:
		var zeroValue T
		return zeroValue, false
	}
}

// Peek returns the top element without removing it
func (s *Stack[T]) Peek() (T, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	select {
	case value := <-s.ch:
		s.ch <- value // push it back to simulate peek
		return value, true
	default:
		var zeroValue T
		return zeroValue, false
	}
}

// Size returns the current size of the stack
func (s *Stack[T]) Size() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return len(s.ch)
}
