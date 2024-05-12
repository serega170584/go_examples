package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

func main() {
	node := &Node{val: 1}
	node1 := &Node{val: 2}
	node2 := &Node{val: 3}
	node.next = node1
	node1.next = node2
	fmt.Println(nEndRemove(node, 2))

	for node != nil {
		fmt.Println(node)
		node = node.next
	}
}

func nEndRemove(node *Node, n int) *Node {
	if n == 0 {
		return node
	}

	var prev *Node
	var first *Node
	var searched *Node
	var prevSearched *Node
	counter := 0
	for node != nil {
		if counter == n {
			prevSearched = searched
			searched = searched.next
		}

		if counter < n {
			counter++
		}

		if prev == nil {
			first = node
		}

		if counter == n && searched == nil {
			searched = first
		}

		prev = node
		node = node.next
	}

	if searched == nil {
		return first
	} else if prevSearched == nil {
		first = first.next
	} else {
		prevSearched.next = searched.next
	}

	return first
}
