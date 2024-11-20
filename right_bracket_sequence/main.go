package main

import (
	"bufio"
	"fmt"
	"os"
)

type Stack struct {
	list       []rune
	startIndex int
}

func NewStack() *Stack {
	s := &Stack{}
	s.startIndex = -1
	return s
}

func (s *Stack) push(r rune) {
	s.startIndex++
	if s.startIndex == len(s.list) {
		s.list = append(s.list, r)
		return
	}
	s.list[s.startIndex] = r
}

func (s *Stack) pop() *rune {
	if s.startIndex == -1 {
		return nil
	}
	startIndex := s.startIndex
	s.startIndex--
	return &s.list[startIndex]
}

func main() {
	const maxCapacity = 3 * 1024 * 1024
	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	scanner.Scan()
	sequence := scanner.Text()

	s := NewStack()

	cnt := 0

	for _, r := range sequence {
		if r == []rune(")")[0] {
			rp := s.pop()
			if rp == nil {
				fmt.Println("no")
				return
			}
			orn := *rp
			if orn != []rune("(")[0] {
				fmt.Println("no")
				return
			}
			cnt--
			continue
		}
		if r == []rune("]")[0] {
			rp := s.pop()
			if rp == nil {
				fmt.Println("no")
				return
			}
			orn := *rp
			if orn != []rune("[")[0] {
				fmt.Println("no")
				return
			}
			cnt--
			continue
		}
		if r == []rune("}")[0] {
			rp := s.pop()
			if rp == nil {
				fmt.Println("no")
				return
			}
			orn := *rp
			if orn != []rune("{")[0] {
				fmt.Println("no")
				return
			}
			cnt--
			continue
		}

		cnt++
		s.push(r)
	}

	if cnt == 0 {
		fmt.Println("yes")
		return
	}

	fmt.Println("no")
}
