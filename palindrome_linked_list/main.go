package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

type Stack struct {
	list []*Node
}

func NewStack() *Stack {
	return &Stack{list: make([]*Node, 0)}
}

func (s *Stack) push(el *Node) {
	s.list = append(s.list, el)
}

func (s *Stack) pop() *Node {
	if len(s.list) == 0 {
		return nil
	}

	el := s.list[len(s.list)-1]
	s.list = s.list[0 : len(s.list)-1]

	return el
}

func main() {
	n1 := &Node{val: 1}
	n2 := &Node{val: 2}
	n1.next = n2
	n3 := &Node{val: 3}
	n2.next = n3
	n4 := &Node{val: 2}
	n3.next = n4
	n5 := &Node{val: 1}
	n4.next = n5

	fmt.Println(isPalindrome(n1))
}

func isPalindrome(head *Node) bool {
	fast := head
	slow := head
	st := NewStack()

	for fast != nil && fast.next != nil {
		st.push(slow)
		slow = slow.next
		fast = fast.next.next
	}

	if fast != nil {
		slow = slow.next
	}

	el := st.pop()
	for slow != nil {
		if slow.val != el.val {
			return false
		}
		slow = slow.next
		el = st.pop()
	}

	return true

}
