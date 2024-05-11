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
	var copyNode *Node
	var prevCopyNode *Node
	linkedNodes := make(map[*Node]*Node)
	var first *Node
	for node != nil {
		if _, ok := linkedNodes[node.random]; !ok {
			linkedNodes[node.random] = &Node{}
		}

		if _, ok := linkedNodes[node]; !ok {
			copyNode = &Node{val: node.val}
			linkedNodes[node] = copyNode
		} else {
			copyNode = linkedNodes[node]
			copyNode.val = node.val
		}

		if prevCopyNode == nil {
			first = copyNode
		} else {
			prevCopyNode.next = copyNode
		}

		copyNode.random = linkedNodes[node.random]

		prevCopyNode = copyNode
		node = node.next
	}

	return first
}
