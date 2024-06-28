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
	n5 := &Node{val: 2}
	n4.next = n5

	fmt.Println(isPalindrome(n1))
}

func isPalindrome(n *Node) bool {
	fast := n
	slow := n
	half := make([]*Node, 0)
	for fast != nil && fast.next != nil {
		half = append(half, slow)
		slow = slow.next
		fast = fast.next.next
	}

	if fast != nil {
		slow = slow.next
	}

	p := len(half) - 1
	for slow != nil {
		if slow.val != half[p].val {
			return false
		}
		slow = slow.next
		p--
	}

	return true
}
