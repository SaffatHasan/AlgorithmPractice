package model

import "fmt"

type Stack struct {
	data []int
}

func (s *Stack) Pop() (result int, err error) {
	if len(s.data) == 0 {
		err = fmt.Errorf("tried to pop from empty stack")
		return
	}
	result = s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return
}

func (s *Stack) Push(x int) {
	s.data = append(s.data, x)
}

func (s *Stack) Peek() int {
	return s.data[s.Len()-1]
}

func (s *Stack) Len() int {
	return len(s.data)
}
