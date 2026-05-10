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

	k := 2

	i := 1
	var mainHead *ListNode

	if k == 1 {
		node = head
		for node != nil {
			fmt.Println(node)
			node = node.Next
		}
		return
	}

	node = head
	tail := node
	head = nil

	for node != nil {
		next := node.Next
		node.Next = head
		head = node

		if i == k {
			head = nil
			mainHead = node
			i = 1
			node = next
			break
		}

		node = next
		i++
	}

	for node != nil {
		next := node.Next
		node.Next = head
		head = node

		if i == 1 {
			tail.Next = node
		}

		if i == k {
			head = nil
			tmpTail := tail.Next
			tail.Next = node
			tail = tmpTail
			i = 0
		}

		node = next
		i++
	}

	node = mainHead
	for node != nil {
		fmt.Println(node)
		node = node.Next
	}
}
