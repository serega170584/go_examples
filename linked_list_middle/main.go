package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

func main() {
	head := &Node{val: 1}
	middle := &Node{val: 2}
	middle1 := &Node{val: 3}
	last := &Node{val: 4}
	head.next = middle
	middle.next = middle1
	middle1.next = last
	fmt.Println(getMiddle(head))
}

func getMiddle(node *Node) *Node {
	runner := node
	for runner.next != nil && runner.next.next != nil {
		runner = runner.next.next
		node = node.next
	}

	return node
}
