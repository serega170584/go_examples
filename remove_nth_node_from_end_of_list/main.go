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

	n := 5
	i := 0
	node = head
	for i < n {
		i++
		node = node.Next
	}

	search := head
	var prev *ListNode
	for node != nil {
		node = node.Next
		prev = search
		search = search.Next
	}

	if prev == nil {
		head = search.Next
	} else {
		prev.Next = search.Next
	}

	node = head
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}
