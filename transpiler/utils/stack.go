package utils

type Stack[T any] struct {
	stack []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(v T) {
	s.stack = append(s.stack, v)
}

func (s *Stack[T]) Pop() T {
	last := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]

	return last
}

func (s *Stack[T]) Top() T {
	return s.stack[len(s.stack)-1]
}
