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
	node4 := &Node{val: 4}
	node3.next = node4
	node5 := &Node{val: 5}
	node4.next = node5

	node := find(node1, 3)
	fmt.Println(node.val)
}

func find(node *Node, k int) *Node {
	_, node = getKthToLast(node, k)
	return node
}

func getKthToLast(node *Node, k int) (int, *Node) {
	if node == nil {
		return 0, nil
	}

	kth, n := getKthToLast(node.next, k)
	if kth == k {
		return kth, n
	}
	kth++
	if kth == k {
		return kth, node
	}

	return kth, nil
}
