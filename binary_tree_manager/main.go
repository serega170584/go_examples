package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	fmt.Println("Enter binary tree count")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	tree := NewBinarySearchTree(n)

	fmt.Println("Enter root value")
	scanner.Scan()
	rootX, _ := strconv.Atoi(scanner.Text())

	root := tree.createNode(rootX)

	for i := 1; i < n; i++ {
		scanner.Scan()
		x, _ := strconv.Atoi(scanner.Text())
		tree.add(root, x)
	}

	fmt.Println(tree)

	fmt.Println("Enter find number")
	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())

	xInd := tree.find(root, x)
	fmt.Println("Got index", xInd)

	leftInd := tree.nodes[xInd].left
	fmt.Println("Got left number", tree.nodes[leftInd].key)

	rightInd := tree.nodes[xInd].right
	fmt.Println("Got right number", tree.nodes[rightInd].key)
}

type Node struct {
	left  int
	right int
	next  int
	key   int
}

type BinarySearchTree struct {
	nodes []*Node
	cnt   int
	first int
}

func NewBinarySearchTree(n int) *BinarySearchTree {
	nodes := make([]*Node, n)
	for i := 0; i < n; i++ {
		nodes[i] = &Node{left: -1, right: -1, next: i + 1}
	}
	return &BinarySearchTree{nodes: nodes, cnt: n}
}

func (tree *BinarySearchTree) find(root, x int) int {
	rootNode := tree.nodes[root]
	key := rootNode.key
	if x < key {
		left := rootNode.left
		if left == -1 {
			return -1
		} else {
			return tree.find(left, x)
		}
	}

	if x > key {
		right := rootNode.right
		if right == -1 {
			return -1
		} else {
			return tree.find(right, x)
		}
	}

	return root
}

func (tree *BinarySearchTree) newNode() int {
	i := tree.first
	tree.first = tree.nodes[i].next
	return i
}

func (tree *BinarySearchTree) createNode(key int) int {
	i := tree.newNode()
	node := tree.nodes[i]
	node.key = key
	node.left = -1
	node.right = -1
	return i
}

func (tree *BinarySearchTree) add(root int, x int) {
	rootNode := tree.nodes[root]
	key := rootNode.key
	if x < key {
		if rootNode.left == -1 {
			i := tree.createNode(x)
			tree.nodes[root].left = i
		} else {
			tree.add(rootNode.left, x)
		}
		return
	}

	if x > key {
		if rootNode.right == -1 {
			i := tree.createNode(x)
			tree.nodes[root].right = i
		} else {
			tree.add(rootNode.right, x)
		}
		return
	}
}
