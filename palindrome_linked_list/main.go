package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

type Stack struct {
	list []int
}

func (st *Stack) Push(v int) {
	st.list = append(st.list, v)
}

func (st *Stack) Pop() *int {
	if st.list == nil {
		return nil
	}
	v := st.list[len(st.list)-1]
	st.list = st.list[0 : len(st.list)-1]
	return &v
}

func NewStack() *Stack {
	return &Stack{}
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

	stack := NewStack()
	for fast != nil && fast.next != nil {
		stack.Push(slow.val)
		slow = slow.next
		fast = fast.next.next
	}

	if fast != nil {
		slow = slow.next
	}

	for slow != nil {
		top := stack.Pop()
		if top == nil {
			return false
		}
		if *top != slow.val {
			return false
		}
		slow = slow.next
	}
	return true
}
