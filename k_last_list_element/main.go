package main

import "fmt"

type Node struct {
	next *Node
	val  int
}

func main() {
	head := &Node{val: 1}
	middle := &Node{val: 2}
	last := &Node{val: 3}
	head.next = middle
	middle.next = last
	fmt.Println(getLastElement(head, 1, new(int)))
}

func getLastElement(node *Node, k int, counter *int) *Node {
	if node.next != nil {
		res := getLastElement(node.next, k, counter)
		if res != nil {
			return res
		}
		if *counter == k {
			return node
		}
	}
	*counter++
	return nil
}
