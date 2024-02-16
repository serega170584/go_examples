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
	list     []int
}

func NewStack(n int) *Stack {
	list := make([]int, n)
	return &Stack{capacity: n, list: list}
}

func (s *Stack) push(el int) {
	i := s.length
	s.list[i] = el
	s.length++
}

func (s *Stack) pop() int {
	i := s.length - 1
	s.length--
	return s.list[i]
}

func factorial(n int) int {
	stack := NewStack(n)
	for i := n; i > 0; i-- {
		stack.push(i)
	}

	res := 1
	for stack.length > 0 {
		res *= stack.pop()
	}

	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	fmt.Println("Enter number")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Got factorial: ", factorial(n))
}
