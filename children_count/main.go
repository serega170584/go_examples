package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Node struct {
	parent      string
	children    []string
	childrenCnt int
	level       int
}

func (n *Node) AddChild(child string) {
	n.children = append(n.children, child)
}

func (n *Node) Children() []string {
	return n.children
}

func (n *Node) SetParent(parent string) {
	n.parent = parent
}

func (n *Node) SetChildrenCnt(childrenCnt int) {
	n.childrenCnt = childrenCnt
}

func (n *Node) ChildrenCnt() int {
	return n.childrenCnt
}

func (n *Node) Parent() string {
	return n.parent
}

func (n *Node) SetLevel(level int) {
	n.level = level
}

func (n *Node) Level() int {
	return n.level
}

func NewNode() *Node {
	return &Node{}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	nodes := make(map[string]*Node, n-1)
	list := make([]string, 0, n-1)
	for i := 0; i < n-1; i++ {
		scanner.Scan()
		relation := strings.Split(scanner.Text(), " ")
		child, parent := relation[0], relation[1]
		var childNode *Node
		if _, ok := nodes[child]; ok {
			childNode = nodes[child]
		} else {
			childNode = NewNode()
			nodes[child] = childNode
			list = append(list, child)
		}

		var parentNode *Node
		if _, ok := nodes[parent]; ok {
			parentNode = nodes[parent]
		} else {
			parentNode = NewNode()
			nodes[parent] = parentNode
			list = append(list, parent)
		}

		nodes[child].SetParent(parent)
		nodes[parent].AddChild(child)
	}

	var root *Node
	for _, v := range nodes {
		if v.Parent() == "" {
			root = v
		}
	}

	children := make([]*Node, 0)
	children = append(children, root)
	levelNodes := make([]*Node, 0, n-1)
	levelNodes = append(levelNodes, root)
	level := 1
	var current []*Node
	for children != nil {
		for _, nodeChildren := range children {
			for _, child := range nodeChildren.Children() {
				nodes[child].SetLevel(level)
				current = append(current, nodes[child])
				levelNodes = append(levelNodes, nodes[child])
			}
		}
		level++
		children = current
		current = nil
	}

	for i := len(levelNodes) - 1; i >= 0; i-- {
		node := levelNodes[i]
		for _, child := range node.Children() {
			node.SetChildrenCnt(node.ChildrenCnt() + nodes[child].ChildrenCnt() + 1)
		}
	}

	slices.Sort(list)

	for _, v := range list {
		fmt.Println(v, nodes[v].ChildrenCnt())
	}
}
