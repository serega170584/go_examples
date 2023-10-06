package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	val      int
	nextNode *Node
}

type LinkedList struct {
	firstNode *Node
	curNode   *Node
}

func (list *LinkedList) add(node *Node) {
	list.curNode.nextNode = node
	list.curNode = node
}

func (list *LinkedList) next() *Node {
	list.curNode = list.curNode.nextNode
	return list.curNode
}

func (list *LinkedList) isEmpty() bool {
	return list.curNode == nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	cnt, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	list := &LinkedList{}
	node := &Node{}
	node.val, _ = strconv.Atoi(scanner.Text())
	list.firstNode, list.curNode = node, node

	middle := make([]*Node, 2)
	middle[0] = node

	for i := 1; i < cnt; i++ {
		scanner.Scan()
		node := &Node{}
		node.val, _ = strconv.Atoi(scanner.Text())
		list.add(node)

		if i%2 == 1 {
			middle[1] = middle[0].nextNode
		}

		if i%2 == 0 {
			middle[0] = middle[1]
			middle[1] = nil
		}
	}

	middleNode := middle[0]
	if middle[1] != nil {
		middleNode = middle[1]
	}

	next := middleNode.nextNode
	middleNode.nextNode = nil

	halfList := swap(middleNode, next)

	halfList.curNode = halfList.firstNode
	list.curNode = list.firstNode

	for !halfList.isEmpty() {
		halfCurNode := halfList.curNode
		curNode := list.curNode

		if halfCurNode.val != curNode.val {
			fmt.Println("NO")
			return
		}

		halfList.next()
		list.next()
	}

	fmt.Println("YES")

}

func swap(node, next *Node) *LinkedList {
	if next == nil {
		list := &LinkedList{curNode: node, firstNode: node}
		return list
	}

	if next.nextNode == nil {
		next.nextNode = node
		list := &LinkedList{curNode: next, firstNode: next}
		return list
	}

	list := swap(next, next.nextNode)

	next.nextNode = node

	return list
}
