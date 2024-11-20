package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Stack struct {
	list []int
}

func NewStack(capacity int) *Stack {
	stack := &Stack{list: make([]int, 0, capacity)}
	return stack
}

func (s *Stack) Push(n int) {
	s.list = append(s.list, n)
}

func (s *Stack) Pop() *int {
	if len(s.list) == 0 {
		return nil
	}
	item := s.list[len(s.list)-1]
	s.list = s.list[:len(s.list)-1]

	return &item
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	list := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		list[i], _ = strconv.Atoi(scanner.Text())
	}

	left := NewStack(k)
	right := NewStack(k)

	for i := 0; i < k; i++ {
		right.Push(list[i])
	}

	item := right.Pop()
	minVal := *item
	left.Push(*item)
	item = right.Pop()
	for item != nil {
		if *item < minVal {
			minVal = *item
		}
		left.Push(minVal)
		item = right.Pop()
	}

	minWindowList := make([]any, 0, n-k)
	minWindowList = append(minWindowList, *left.Pop())

	ind := k
	for ind != n {
		item = left.Pop()
		if item != nil {
			rightItem := list[ind]
			right.Push(rightItem)
			v := *item
			minRight := rightItem
			if minRight < v {
				v = minRight
			}
			minWindowList = append(minWindowList, v)
			item = left.Pop()
			ind++

			for item != nil && ind != len(list) {
				rightItem = list[ind]
				if rightItem < minRight {
					minRight = rightItem
				}
				v = *item
				if minRight < v {
					v = minRight
				}
				minWindowList = append(minWindowList, v)
				ind++
				item = left.Pop()
				right.Push(rightItem)
			}
		}

		if ind == n {
			break
		}

		minVal = list[ind]
		ind++
		left.Push(minVal)
		item = right.Pop()
		for item != nil {
			if *item < minVal {
				minVal = *item
			}
			left.Push(minVal)
			item = right.Pop()
		}

		minWindowList = append(minWindowList, *left.Pop())
	}

	for _, v := range minWindowList {
		fmt.Println(v)
	}
}
