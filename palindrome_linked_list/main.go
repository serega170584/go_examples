package main

import "fmt"

type Node struct {
	next *Node
	val  int
}

func main() {
	node := &Node{val: 1}
	next := node
	node = &Node{val: 1}
	node.next = next
	//next = node
	//node = &Node{val: 3}
	//node.next = next
	//next = node
	//node = &Node{val: 2}
	//node.next = next
	//next = node
	//node = &Node{val: 1}
	//node.next = next

	fmt.Println(isPalindrom(node))
}

func isPalindrom(node *Node) bool {
	var firstList *Node
	var secondList *Node
	var middle *Node
	var nextFirst *Node
	counter := 0
	for node != nil {
		counter++
		if counter == 1 {
			middle = node
			node = node.next
			continue
		}

		if counter%2 == 0 {
			nextFirst = firstList
			firstList = middle
			middle = middle.next
			firstList.next = nextFirst
			secondList = middle
		} else {
			secondList = secondList.next
		}

		node = node.next
	}

	counter /= 2
	for i := 0; i < counter; i++ {
		if firstList.val != secondList.val {
			return false
		}
		firstList = firstList.next
		secondList = secondList.next
	}

	return true
}
