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
	fmt.Println(getMiddle(n1))
}

func getMiddle(node *Node) *Node {
	slow := node
	fast := node
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
	}

	return slow

}
