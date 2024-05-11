package main

import "fmt"

type Node struct {
	val    int
	next   *Node
	random *Node
}

func main() {
	node3 := &Node{
		val: 3,
	}
	node2 := &Node{
		val:  2,
		next: node3,
	}
	node1 := &Node{
		val:  1,
		next: node2,
	}
	node3.random = node1
	node2.random = node3
	node1.random = node3

	node := copyList(node1)
	for node != nil {
		fmt.Println(node)
		node = node.next
	}
}

func copyList(node *Node) *Node {
	var prev *Node
	var first *Node
	var copyNode *Node
	linkedNodes := make(map[*Node]*Node)
	for node != nil {
		if _, ok := linkedNodes[node]; ok {
			copyNode = linkedNodes[node]
			copyNode.val = node.val
		} else {
			copyNode = &Node{val: node.val}
			linkedNodes[node] = copyNode
		}

		if _, ok := linkedNodes[node.random]; ok {
			copyNode.random = linkedNodes[node.random]
		} else {
			copyNode.random = &Node{}
			linkedNodes[node.random] = copyNode.random
		}

		if prev == nil {
			first = copyNode
		} else {
			prev.next = copyNode
		}

		prev = copyNode
		node = node.next
	}

	return first
}
