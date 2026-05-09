package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node := &ListNode{
		Val:  0,
		Next: nil,
	}

	head := node
	for i := 1; i < 5; i++ {
		prev := node
		node = &ListNode{
			Val: i,
		}
		prev.Next = node
	}

	n := 1
	i := 0
	node = head
	for node != nil {
		if i == n {
			break
		}
		node = node.Next
		i++
	}

	mainHead := head
	var prev *ListNode
	for node != nil {
		prev = head
		head = head.Next
		node = node.Next
	}

	prev.Next = head.Next

	node = mainHead
	for node != nil {
		fmt.Println(node)
		node = node.Next
	}
}
