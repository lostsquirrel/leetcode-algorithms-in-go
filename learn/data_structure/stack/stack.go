package stack

import (
	"errors"
)

type IntStack struct {
	size int
	top  int
	data []int
}

func NewIntStack(size int) *IntStack {
	s := &IntStack{
		size: size,
		top:  -1,
		data: make([]int, size),
	}
	return s
}

func (s *IntStack) Push(n int) error {
	if s.top < s.size-1 {
		s.top += 1
		s.data[s.top] = n
		return nil
	}
	return errors.New("stack overflow")
}

func (s *IntStack) Pop() (int, error) {
	r, e := s.Top()
	if e != nil {
		return 0, e
	}
	s.data[s.top] = 0
	s.top -= 1
	return r, nil
}

func (s *IntStack) Top() (int, error) {
	if s.top < 0 {
		return 0, errors.New("stack empty")
	}
	r := s.data[s.top]
	return r, nil
}

func (s *IntStack) Empty() bool {
	return s.top == -1
}
