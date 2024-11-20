package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	list []int
}

func NewStack(n int) *Stack {
	list := make([]int, 0, n)
	return &Stack{list: list}
}

func (s *Stack) push(item int) {
	s.list = append(s.list, item)
}

func (s *Stack) pop() *int {
	if len(s.list) == 0 {
		return nil
	}

	item := s.list[len(s.list)-1]
	s.list = s.list[0 : len(s.list)-1]
	return &item
}

func main() {
	const maxCapacity = 3 * 1024 * 1024
	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	scanner.Scan()
	t := scanner.Text()
	list := strings.Split(t, " ")

	resSt := NewStack(len(list))

	for _, v := range list {
		if v == "" {
			break
		}
		if v != "+" && v != "-" && v != "*" {
			item, _ := strconv.Atoi(v)
			resSt.push(item)
			continue
		}

		b := *resSt.pop()
		a := *resSt.pop()
		res := a + b

		if v == "*" {
			res = a * b
		}

		if v == "-" {
			res = a - b
		}

		resSt.push(res)
	}

	fmt.Println(*resSt.pop())
}
