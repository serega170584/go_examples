package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Memory struct {
	cnt   int
	nodes []*Node
	first int
}

type Node struct {
	leftKey  int
	rightKey int
	key      int
}

func NewMemory(cnt int) *Memory {
	nodes := make([]*Node, cnt)
	for i := range nodes {
		node := &Node{leftKey: i + 1}
		nodes[i] = node
	}
	return &Memory{cnt: cnt, nodes: nodes, first: 0}
}

func (m *Memory) addNode() {
	node := m.nodes[m.first]
	m.first = node.leftKey
}

func (m *Memory) delNode(i int) {
	m.nodes[i].leftKey = m.first
	m.first = i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	fmt.Println("Enter memory length")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	m := NewMemory(n)
	fmt.Println(m)

	m.addNode()
	fmt.Println(m)

	m.addNode()
	fmt.Println(m)

	m.addNode()
	fmt.Println(m)

	m.delNode(1)
	fmt.Println(m)
}
