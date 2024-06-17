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
	var middle *Node
	var firstHead *Node
	var secondHead *Node
	var prevFirstHead *Node
	for node != nil {
		counter++
		if counter == 1 {
			middle = node
			node = node.next
			continue
		}

		if counter%2 == 0 {
			prevFirstHead = firstHead
			firstHead = middle
			middle = middle.next
			firstHead.next = prevFirstHead
			secondHead = middle
		} else {
			secondHead = middle.next
		}

		node = node.next
	}

	counter /= 2
	for i := 0; i < counter; i++ {
		if firstHead.val != secondHead.val {
			return false
		}
		firstHead = firstHead.next
		secondHead = secondHead.next
	}

	return true
}
