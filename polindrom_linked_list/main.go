package main

import "fmt"

type Node struct {
	next *Node
	val  int
}

func main() {
	node3 := &Node{val: 1}
	node2 := &Node{val: 2, next: node3}
	node1 := &Node{val: 2, next: node2}
	fmt.Println(isPolindrom(node1))
}

func isPolindrom(node *Node) bool {
	var middle *Node
	cnt := 0
	var prevFirst *Node
	var first *Node
	var secondHalfFirst *Node
	for node != nil {
		cnt++
		if cnt == 1 {
			middle = node
		}

		if cnt%2 == 0 {
			prevFirst = first
			first = middle
			middle = middle.next
			first.next = prevFirst
			secondHalfFirst = middle
		} else if cnt != 1 {
			secondHalfFirst = secondHalfFirst.next
		}

		node = node.next
	}

	for first != nil && secondHalfFirst != nil {
		if first.val != secondHalfFirst.val {
			return false
		}
		first = first.next
		secondHalfFirst = secondHalfFirst.next
	}

	return true
}
