package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

func main() {
	node1 := &Node{val: 1}
	node2 := &Node{val: 2}
	node1.next = node2
	node3 := &Node{val: 3}
	node2.next = node3
	node4 := &Node{val: 7}
	node3.next = node4
	node5 := &Node{val: 8}
	node4.next = node5
	node6 := &Node{val: 9}
	node5.next = node6
	node7 := &Node{val: 4}
	node8 := &Node{val: 5}
	node7.next = node8
	node9 := &Node{val: 6}
	node8.next = node9
	node9.next = node4
	node := getIntersection(node1, node7)
	for node != nil {
		fmt.Println(node.val)
		node = node.next
	}
}

func getIntersection(list1 *Node, list2 *Node) *Node {
	s1, tail1 := getSizeAndTail(list1)
	s2, tail2 := getSizeAndTail(list2)
	var shorter *Node
	var longer *Node
	if tail1 != tail2 {
		return nil
	}
	if s1 < s2 {
		shorter = list1
		longer = list2
	} else {
		shorter = list2
		longer = list1
	}
	diff := s1 - s2
	if diff < 0 {
		diff = -diff
	}
	longer = getKthNode(longer, diff)

	for shorter != longer {
		shorter = shorter.next
		longer = longer.next
	}

	return longer
}

func getSizeAndTail(node *Node) (int, *Node) {
	counter := 0
	var tail *Node
	for node != nil {
		counter++
		tail = node
		node = node.next
	}

	return counter, tail
}

func getKthNode(node *Node, k int) *Node {
	counter := 0
	for counter != k {
		node = node.next
		counter++
	}

	return node
}
