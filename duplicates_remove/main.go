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
	n4 := &Node{val: 4}
	n3.next = n4
	n5 := &Node{val: 2}
	n4.next = n5
	n6 := &Node{val: 2}
	n5.next = n6

	removeDuplicates(n1)
	node := n1
	for node != nil {
		fmt.Println(node.val)
		node = node.next
	}
}

func removeDuplicates(node *Node) *Node {
	head := node
	for node != nil {
		runner := node
		for runner.next != nil {
			if runner.next.val == node.val {
				runner.next = runner.next.next
			} else {
				runner = runner.next
			}
		}
		node = node.next
	}

	return head
}
