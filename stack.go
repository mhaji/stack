package stack

import (
	"sync"
	"errors"
)

var EmptyStackErr = errors.New("Stack is empty")

type stack struct {
	data []interface{}
	sync.Mutex
}

func NewStack() *stack {
	return new(stack)
}

func (s *stack) Push(elem interface{}) {
	s.Lock()
	defer s.Unlock()

	s.data = append(s.data, elem)
}

func (s *stack) Pop() (interface{}, error) {
	s.Lock()
	defer s.Unlock()

	if s.Size() == 0 {
		return nil, EmptyStackErr
	}

	ret := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return ret, nil
}

func (s *stack) Size() int {
	return len(s.data)
}