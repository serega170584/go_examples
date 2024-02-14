package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type node struct {
	key   int
	freq  int
	left  int
	right int
	next  int
	index int
}

type HaffmanTree struct {
	nodes     []*node
	first     int
	root      int
	heap      *BinaryHeap
	cnt       int
	codeTable map[int]string
}

type BinaryHeap struct {
	nodes []*node
	cnt   int
	last  int
}

func NewBinaryHeap(cnt int) *BinaryHeap {
	heap := &BinaryHeap{cnt: cnt}
	heap.nodes = make([]*node, cnt)
	return heap
}

func (heap *BinaryHeap) add(node *node) {
	last := heap.last
	heap.nodes[last] = node
	for last != 0 {
		parent := (last - 1) / 2
		if heap.nodes[last].freq < heap.nodes[parent].freq {
			heap.nodes[last], heap.nodes[parent] = heap.nodes[parent], heap.nodes[last]
		}
		last = parent
	}
	heap.last++
}

func (heap *BinaryHeap) popMin() *node {
	last := heap.last
	node := heap.nodes[0]
	heap.nodes[0] = heap.nodes[last-1]
	last--
	heap.last = last
	current := 0
	for current < last-1 {
		left := 2*current + 1
		right := 2*current + 2

		minInd := current

		if left < last && heap.nodes[minInd].freq > heap.nodes[left].freq {
			minInd = left
		}

		if right < last && heap.nodes[minInd].freq > heap.nodes[right].freq {
			minInd = right
		}

		if minInd == current {
			break
		}

		heap.nodes[minInd], heap.nodes[current] = heap.nodes[current], heap.nodes[minInd]
		current = minInd
	}

	return node
}

func NewHaffmanTree(cnt int) *HaffmanTree {
	listCnt := 2*cnt - 1
	nodes := make([]*node, listCnt)
	for i := 0; i < listCnt; i++ {
		nodes[i] = &node{next: i + 1, index: i}
	}
	heap := NewBinaryHeap(cnt)
	codeTable := make(map[int]string, cnt)
	return &HaffmanTree{nodes: nodes, heap: heap, cnt: listCnt, codeTable: codeTable}
}

func (tree *HaffmanTree) createNode(x int, freq int, left int, right int) int {
	first := tree.first
	node := tree.nodes[tree.first]
	node.key = x
	node.freq = freq
	node.left = left
	node.right = right
	node.index = first
	tree.first = node.next
	return first
}

func (tree *HaffmanTree) build() {
	for tree.heap.last > 1 {
		left := tree.heap.popMin()
		right := tree.heap.popMin()
		ind := tree.createNode(-1, left.freq+right.freq, left.index, right.index)
		tree.heap.add(tree.nodes[ind])
		tree.root = ind
	}
}

func (tree *HaffmanTree) find(x int) int {
	for i := 0; i < tree.cnt; i++ {
		if tree.nodes[i].key == x {
			return i
		}
	}

	return -1
}

func (tree *HaffmanTree) buildCodeTable(ind int, code string) {
	el := tree.nodes[ind]
	if el.left == -1 && el.right == -1 {
		tree.codeTable[el.key] = code
	} else {
		left := el.left
		if left != -1 {
			tree.buildCodeTable(left, code+"0")
		}
		right := el.right
		if right != -1 {
			tree.buildCodeTable(right, code+"1")
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	fmt.Println("Enter dictionary size")
	scanner.Scan()
	dictSize, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter list count")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter list")
	tree := NewHaffmanTree(dictSize)
	for i := 0; i < n; i++ {
		scanner.Scan()
		el, _ := strconv.Atoi(scanner.Text())
		ind := tree.find(el)
		if ind == -1 {
			ind = tree.createNode(el, 0, -1, -1)
		}
		tree.nodes[ind].freq++
	}

	for i := 0; i < dictSize; i++ {
		tree.heap.add(tree.nodes[i])
	}

	tree.build()
	tree.buildCodeTable(tree.root, "")
	fmt.Println(tree.codeTable)
}
