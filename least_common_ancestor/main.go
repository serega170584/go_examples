package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	parent   string
	children []string
}

func CreateNode() *Node {
	return &Node{}
}

func (n *Node) setParent(parent string) {
	n.parent = parent
}

func (n *Node) Parent() string {
	return n.parent
}

func (n *Node) addChild(child string) {
	n.children = append(n.children, child)
}

func (n *Node) Children() []string {
	return n.children
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	nodes := make(map[string]*Node, n-1)
	for i := 0; i < n-1; i++ {
		scanner.Scan()
		relation := strings.Split(scanner.Text(), " ")
		child, parent := relation[0], relation[1]
		var childNode *Node
		if _, ok := nodes[child]; ok {
			childNode = nodes[child]
		} else {
			childNode = CreateNode()
			nodes[child] = childNode
		}
		childNode.setParent(parent)

		var parentNode *Node
		if _, ok := nodes[parent]; ok {
			parentNode = nodes[parent]
		} else {
			parentNode = CreateNode()
			nodes[parent] = parentNode
		}
		parentNode.addChild(child)
	}

	requests := make([][2]string, 0)
	for scanner.Scan() {
		relation := strings.Split(scanner.Text(), " ")
		requests = append(requests, [2]string{relation[0], relation[1]})
	}

	for _, request := range requests {
		parents1 := make([]string, 0)
		parents1 = append(parents1, request[0])
		parent := nodes[request[0]].Parent()
		for parent != "" {
			parents1 = append(parents1, parent)
			parent = nodes[parent].Parent()
		}

		parents2 := make([]string, 0)
		parents2 = append(parents1, request[1])
		parent = nodes[request[1]].Parent()
		for parent != "" {
			parents2 = append(parents2, parent)
			parent = nodes[parent].Parent()
		}

		prev := parents1[len(parents1)-1]
		parent1P := len(parents1) - 2
		parent2P := len(parents2) - 2
		for parent1P >= 0 && parent2P >= 0 && parents1[parent1P] == parents2[parent2P] {
			prev = parents1[parent1P]
			parent1P--
			parent2P--
		}

		fmt.Println(prev)
	}

}
