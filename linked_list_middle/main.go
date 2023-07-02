package main

import (
	"fmt"
)

func main() {
	data := [6]int{1, 2, 3, 4, 5, 6}

	list := &linkedList{}
	for _, val := range data {
		el := &item{
			value: val,
		}
		list.add(el)
	}
	fmt.Println(list.getMiddle().value)
}

type item struct {
	value int
	next  *item
}

type linkedList struct {
	root     *item
	middle   *item
	current  *item
	curIndex int
}

func (list *linkedList) add(el *item) {
	if list.root == nil {
		list.root = el
		list.middle = el
		list.current = list.root
	} else {
		list.current.next = el
		list.current = el
	}

	if list.curIndex%2 == 1 {
		list.middle = list.middle.next
	}

	list.curIndex++
}

func (list *linkedList) getMiddle() *item {
	return list.middle
}

func (list *linkedList) getCurrent() *item {
	return list.current
}
