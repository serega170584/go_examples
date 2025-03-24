package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

func main() {
	n1 := &Node{val: 1}
	n2 := &Node{val: 2}
	n1.next = n2
	n3 := &Node{val: 3}
	n2.next = n3
	n4 := &Node{val: 4}
	n3.next = n4
	n5 := &Node{val: 5}
	n4.next = n5
	n6 := &Node{val: 6}
	n5.next = n6
	n7 := &Node{val: 7}
	n6.next = n7
	n8 := &Node{val: 8}
	n7.next = n8
	n9 := &Node{val: 9}
	n8.next = n9
	n9.next = n4
	fmt.Println(findListLoopBeginning(n1))
}

func findListLoopBeginning(head *Node) *Node {
	slow := head
	fast := head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			break
		}
	}

	if fast == nil || fast.next == nil {
		return nil
	}

	slow = head
	for slow != fast {
		slow = slow.next
		fast = fast.next
	}

	return slow

}
