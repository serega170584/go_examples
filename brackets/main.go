package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Stack struct {
	capacity int
	length   int
	list     []string
}

func NewStack(cnt int) *Stack {
	list := make([]string, cnt)
	return &Stack{capacity: cnt, length: 0, list: list}
}

func (s *Stack) push(el string) {
	curIndex := s.length
	s.list[curIndex] = el
	s.length++
}

func (s *Stack) pop() string {
	s.length--
	curIndex := s.length
	return s.list[curIndex]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	fmt.Println("Enter count")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter sequence")
	scanner.Scan()
	sequence := scanner.Text()
	sequenceRunes := []rune(sequence)
	list := make([]string, n)
	for i := 0; i < n; i++ {
		list[i] = string(sequenceRunes[i])
	}

	fmt.Println("Is brackets string valid?", isRightBracketsSequence(n, list))
}

func isRightBracketsSequence(n int, brackets []string) bool {
	stack := NewStack(n)
	balance := 0
	for _, val := range brackets {
		if val == "{" || val == "(" || val == "[" {
			stack.push(val)
			balance++
		}

		if (val == "}" || val == "]" || val == ")") && stack.length == 0 {
			return false
		}

		if val == "}" {
			el := stack.pop()
			if el != "{" {
				return false
			}
		}

		if val == "]" {
			el := stack.pop()
			if el != "[" {
				return false
			}
		}

		if val == ")" {
			el := stack.pop()
			if el != "(" {
				return false
			}
		}

		if val == "}" || val == "]" || val == ")" {
			balance--
		}
	}

	if balance != 0 {
		return false
	}

	return true
}
