package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	n := 5
	head := []int{1, 2, 3, 4, 5}
	node := &ListNode{Val: head[0]}
	headNode := node
	for i := 1; i < 5; i++ {
		node.Next = &ListNode{Val: head[i]}
		node = node.Next
	}

	node = headNode
	i := 0
	for i < n {
		node = node.Next
		i++
	}

	leftNode := headNode
	var prev *ListNode
	for node != nil {
		node = node.Next
		prev = leftNode
		leftNode = leftNode.Next
	}

	if prev == nil {
		headNode = leftNode.Next
	} else {
		prev.Next = leftNode.Next
	}

	node = headNode
	for node != nil {
		fmt.Println(node)
		node = node.Next
	}
}
