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
	children []string
	parent   string
	level    int
}

func NewNode() *Node {
	return &Node{}
}

func (n *Node) AddChild(name string) {
	n.children = append(n.children, name)
}

func (n *Node) SetParent(parent string) {
	n.parent = parent
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

func (n *Node) Children() []string {
	return n.children
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	nodes := make(map[string]*Node, n-1)
	list := make([]string, 0, n-1)
	for i := 0; i < n-1; i++ {
		scanner.Scan()
		ancestors := strings.Split(scanner.Text(), " ")
		child, parent := ancestors[0], ancestors[1]
		var childNode *Node
		if _, ok := nodes[child]; ok {
			childNode = nodes[child]
		} else {
			childNode = NewNode()
			nodes[child] = childNode
			list = append(list, child)
		}

		childNode.SetParent(parent)

		var parentNode *Node
		if _, ok := nodes[parent]; ok {
			parentNode = nodes[parent]
		} else {
			parentNode = NewNode()
			nodes[parent] = parentNode
			list = append(list, parent)
		}

		parentNode.AddChild(child)
	}

	var zeroAncestor *Node
	for _, v := range nodes {
		if v.Parent() == "" {
			zeroAncestor = v
		}
	}

	zeroAncestor.SetLevel(0)
	level := 1
	children := zeroAncestor.Children()
	currentChildren := make([]string, 0)
	for children != nil {
		for _, v := range children {
			nodes[v].SetLevel(level)
			ancestorChildren := nodes[v].Children()
			for _, ancestorChild := range ancestorChildren {
				currentChildren = append(currentChildren, ancestorChild)
			}
		}
		level++
		children = currentChildren
		currentChildren = nil
	}

	slices.Sort(list)

	for _, v := range list {
		fmt.Println(v, nodes[v].Level())
	}
}
