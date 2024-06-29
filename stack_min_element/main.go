package main

import (
	"fmt"
)

type Stack struct {
	items []int
	min   []int
}

func newStack() *Stack {
	return &Stack{}
}

func (s *Stack) push(v int) {
	s.items = append(s.items, v)
	if len(s.min) == 0 || v <= s.min[len(s.min)-1] {
		s.min = append(s.min, v)
	}
}

func (s *Stack) pop() int {
	el := s.items[len(s.items)-1]
	if el == s.min[len(s.min)-1] {
		s.min = s.min[0 : len(s.min)-1]
	}
	s.items = s.items[0 : len(s.items)-1]
	return el
}

func (s *Stack) getMin() int {
	return s.min[len(s.min)-1]
}

func main() {
	s := newStack()
	s.push(5)
	s.push(4)
	s.push(6)
	s.push(7)
	s.push(7)
	s.push(7)
	s.pop()
	fmt.Println(s.getMin())
	s.pop()
	fmt.Println(s.getMin())
	s.pop()
	fmt.Println(s.getMin())
	s.pop()
	fmt.Println(s.getMin())
	s.pop()
	fmt.Println(s.getMin())
}
