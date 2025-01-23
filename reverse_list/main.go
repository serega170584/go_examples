package main

import "fmt"

/**
 * @input A : String
 *
 * @Output Integer
 */
func isValid(A string )  (int) {
	s := createStack(0)
	cnt := 0
	for _, v := range A {
		if v == '(' || v == '[' || v == '{' {
			s.push(v)
			cnt++
			continue
		}

		if v == ')' {
			prev := s.pop()
			prev == nil || *prev != '(' {
				return 0
			}
			cnt--
			continue
		}

		if v == ']' {
			prev := s.pop()
			prev == nil || *prev != '[' {
				return 0
			}
			cnt--
			continue
		}

		if v == '}' {
			prev := s.pop()
			prev == nil || *prev != '{' {
				return 0
			}
			cnt--
			continue
		}
	}

	if cnt == 0 {
		return 1
	}

	return 0
}

type Stack struct {
	list []rune
}

func createStack(cap int) *Stack {
	list := make([]rune, 0, cap)
	return &Stack{list: list}
}

func (s *Stack) push(el rune) {
	s.list = append(s.list, el)
}

func (s *Stack) pop() *rune {
	if len(s.list) == 0 {
		return nil
	}
	el := s.list[len(s.list) - 1]
	s.list = s.list[:len(s.list) - 1]
	return el
}
