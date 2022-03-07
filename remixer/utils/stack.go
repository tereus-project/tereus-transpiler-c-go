package utils

type Stack struct {
	stack []string
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(v string) {
	s.stack = append(s.stack, v)
}

func (s *Stack) Pop() string {
	last := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]

	return last
}

func (s *Stack) Top() string {
	return s.stack[len(s.stack)-1]
}
