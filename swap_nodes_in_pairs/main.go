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

	var prevTail *ListNode
	var prev *ListNode
	isTail := false

	if head.Next != nil {
		node = head.Next
		head.Next = node.Next
		node.Next = head
		head = node
		prevTail = head.Next
		node = head.Next.Next
	}

	for node != nil {
		if isTail {
			prev.Next = node.Next
			node.Next = prev
			prevTail.Next = node
			prevTail = prev
			node = prev
		}

		prev = node
		node = node.Next
		isTail = !isTail
	}

	node = head
	for node != nil {
		fmt.Println(node)
		node = node.Next
	}
}
