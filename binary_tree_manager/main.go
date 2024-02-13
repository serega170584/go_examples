package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 9 4 1 5 6 7 2 3 10
//
//	     8
//	4         9
//
// 1    5
//
//	2      6
//	   3      7
//	              10
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

	root := tree.createRoot(rootX)

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

	fmt.Println("Got visited tree")
	tree.visit(root)
	fmt.Println()

	fmt.Println("Delete element")
	scanner.Scan()
	el, _ := strconv.Atoi(scanner.Text())

	tree.del(root, root, el)

	fmt.Println("Got visited tree")
	tree.visit(root)
	fmt.Println()

	fmt.Println("Delete element")
	scanner.Scan()
	el, _ = strconv.Atoi(scanner.Text())

	tree.del(tree.root, tree.root, el)

	fmt.Println("Got visited tree")
	tree.visit(tree.root)
	fmt.Println()
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
	root  int
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

func (tree *BinarySearchTree) createRoot(root int) int {
	node := tree.createNode(root)
	tree.root = node
	return node
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

func (tree *BinarySearchTree) visit(root int) {
	left := tree.nodes[root].left
	right := tree.nodes[root].right
	key := tree.nodes[root].key

	if left != -1 {
		tree.visit(left)
	}

	fmt.Print(key, " ")

	if right != -1 {
		tree.visit(right)
	}
}

func (tree *BinarySearchTree) del(parent int, current int, x int) {
	currentNode := tree.nodes[current]

	if x < currentNode.key {
		tree.del(current, currentNode.left, x)
		return
	}

	if x > currentNode.key {
		tree.del(current, currentNode.right, x)
		return
	}

	currentLeft := tree.nodes[current].left
	currentRight := tree.nodes[current].right

	parentCornerLeft := -1
	cornerLeft := -1
	if currentRight != -1 {
		parentCornerLeft, cornerLeft = tree.findCornerLeft(currentRight, currentRight)
	}

	first := tree.first
	tree.nodes[current].left = -1
	tree.nodes[current].right = -1
	tree.nodes[current].next = first
	tree.first = current

	if parentCornerLeft != -1 {
		tree.nodes[parentCornerLeft].left = -1
	}

	if cornerLeft != -1 {
		tree.nodes[cornerLeft].left = currentLeft
	}

	if cornerLeft != -1 && cornerLeft != parentCornerLeft {
		tree.nodes[cornerLeft].right = currentRight
	}

	if tree.nodes[parent].left == current {
		tree.nodes[parent].left = cornerLeft
	}

	if tree.nodes[parent].right == current {
		tree.nodes[parent].right = cornerLeft
	}

	if tree.root == current {
		tree.root = cornerLeft
		if cornerLeft == -1 {
			tree.root = tree.nodes[current].left
		}
	}
}

func (tree *BinarySearchTree) findCornerLeft(parent int, current int) (int, int) {
	if tree.nodes[current].left == -1 {
		return parent, current
	}
	return tree.findCornerLeft(current, tree.nodes[current].left)
}
