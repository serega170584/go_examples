package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	fmt.Println("Enter list length")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter list")
	list := make([]int, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		list[i], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println("Got min right list", minRightValues(n, list))
}

type Stack struct {
	capacity int
	length   int
	list     [][2]int
}

func (s *Stack) push(el [2]int) {
	i := s.length
	s.list[i] = el
	s.length++
}

func (s *Stack) pop() [2]int {
	el := s.list[s.length-1]
	s.length--
	return el
}

func NewStack(n int) *Stack {
	capacity := n
	list := make([][2]int, n)
	return &Stack{capacity: capacity, list: list}
}

func minRightValues(n int, list []int) []int {
	minValues := make([]int, n)
	stack := NewStack(n)
	for i, val := range list {
		for stack.length != 0 {
			el := stack.pop()
			if el[1] > val {
				minValues[el[0]] = val
			} else {
				stack.push(el)
				break
			}
		}
		stack.push([2]int{i, val})
	}
	return minValues
}
