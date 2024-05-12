package main

import "fmt"

type Node struct {
	next *Node
	val  int
}

func main() {
	node := &Node{val: 1}
	next := node
	node = &Node{val: 2}
	node.next = next
	next = node
	node = &Node{val: 3}
	node.next = next
	next = node
	node = &Node{val: 2}
	node.next = next
	next = node
	node = &Node{val: 1}
	node.next = next

	fmt.Println(isPalindrom(node))
}

func isPalindrom(node *Node) bool {
	counter := 0
	var first *Node
	var second *Node
	var middle *Node
	var nextFirst *Node
	for node != nil {
		counter++
		if counter == 1 {
			middle = node
			node = node.next
			continue
		}

		if counter%2 == 0 {
			nextFirst = first
			first = middle
			middle = middle.next
			second = middle
			first.next = nextFirst
		} else {
			second = middle.next
		}

		node = node.next
	}

	counter /= 2
	for i := 0; i < counter; i++ {
		if first.val != second.val {
			return false
		}
		first = first.next
		second = second.next
	}

	return true
}
