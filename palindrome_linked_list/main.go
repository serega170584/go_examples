package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

func main() {
	n1 := &Node{val: 1}
	n2 := &Node{val: 2}
	n1.next = n2
	n3 := &Node{val: 3}
	n2.next = n3
	n4 := &Node{val: 2}
	n3.next = n4
	n5 := &Node{val: 1}
	n4.next = n5

	isPalindrome := true
	slow := n1
	fast := n1

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	if fast != nil {
		slow = slow.next
	}

	var prev *Node = nil
	for slow != nil {
		next := slow.next
		slow.next = prev
		prev = slow
		slow = next
	}

	left := n1
	right := prev
	for right != nil {
		if right.val != left.val {
			isPalindrome = false
			break
		}

		left = left.next
		right = right.next
	}

	fmt.Println(isPalindrome)
}
