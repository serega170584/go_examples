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

	newNode := copyList(node1)
	for newNode != nil {
		fmt.Println(newNode.val)
		fmt.Println(newNode.random.val)
		newNode = newNode.next
	}
}

func copyList(node *Node) *Node {
	var first *Node
	var prev *Node
	var newNode *Node
	links := make(map[*Node]*Node)
	for node != nil {
		if _, ok := links[node]; ok {
			links[node].val = node.val
			newNode = links[node]
		} else {
			newNode = &Node{val: node.val}
			links[node] = newNode
		}

		if _, ok := links[node.random]; ok {
			newNode.random = links[node.random]
		} else {
			newNode.random = &Node{}
			links[node.random] = newNode.random
		}

		if prev == nil {
			first = newNode
		} else {
			prev.next = newNode
		}

		prev = newNode
		node = node.next
	}

	return first
}
