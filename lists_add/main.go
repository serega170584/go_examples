package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

func main() {
	n11 := &Node{val: 6}
	n12 := &Node{val: 7}
	n11.next = n12
	n13 := &Node{val: 5}
	n12.next = n13
	n21 := &Node{val: 6}
	n22 := &Node{val: 3}
	n21.next = n22
	n23 := &Node{val: 6}
	n22.next = n23
	n := addLists(n11, n21, 0)
	for n != nil {
		fmt.Println(n.val)
		n = n.next
	}
}

func addLists(n1 *Node, n2 *Node, carry int) *Node {
	if n1 == nil && n2 == nil && carry == 0 {
		return nil
	}

	res := &Node{}
	v := carry
	if n1 != nil {
		v += n1.val
	}
	if n2 != nil {
		v += n2.val
	}

	res.val = v % 10
	var nn1 *Node
	if n1 != nil {
		nn1 = n1.next
	}
	var nn2 *Node
	if n2 != nil {
		nn2 = n2.next
	}
	res.next = addLists(nn1, nn2, v/10)
	return res
}
