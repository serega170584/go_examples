package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	nodes      []*node
	first      int
	root       int
	heap       *BinaryHeap
	cnt        int
	codeTable  map[int]string
	valueTable map[string]int
	sb         strings.Builder
	code       string
	codeInd    int
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
	valueTable := make(map[string]int, cnt)
	return &HaffmanTree{nodes: nodes, heap: heap, cnt: listCnt, codeTable: codeTable, valueTable: valueTable}
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

func (tree *HaffmanTree) visit(root int) {
	left := tree.nodes[root].left
	if left != -1 {
		tree.sb.WriteString("D")
		tree.visit(left)
	}

	right := tree.nodes[root].right
	if right != -1 {
		tree.sb.WriteString("U")
		tree.visit(right)
	}
}

func (tree *HaffmanTree) generateCodeTable() {
	code := tree.code
	codeRunes := []rune(code)
	tree.visitCodeSym(tree.root, codeRunes, "")
}

func (tree *HaffmanTree) generateValueTable() {
	codeTable := tree.codeTable
	for n, code := range codeTable {
		tree.valueTable[code] = n
	}
}

func (tree *HaffmanTree) visitCodeSym(root int, codeRunes []rune, code string) {
	if tree.nodes[root].left == -1 && tree.nodes[root].right == -1 {
		key := tree.nodes[root].key
		tree.codeTable[key] = code
		return
	}

	if string(codeRunes[tree.codeInd]) == "D" {
		tree.codeInd++
		tree.visitCodeSym(tree.nodes[root].left, codeRunes, code+"0")
	}

	if string(codeRunes[tree.codeInd]) == "U" {
		tree.codeInd++
		tree.visitCodeSym(tree.nodes[root].right, codeRunes, code+"1")
	}
}

func (tree *HaffmanTree) generateListCode(list []int) string {
	var sb strings.Builder
	for _, val := range list {
		sb.WriteString(tree.codeTable[val])
	}
	return sb.String()
}

func (tree *HaffmanTree) generateList(code string, cnt int) []int {
	list := make([]int, cnt)
	ind := 0
	codeRunes := []rune(code)
	symCode := ""
	for _, runeVal := range codeRunes {
		val := string(runeVal)
		symCode += val
		if v, ok := tree.valueTable[symCode]; ok {
			list[ind] = v
			symCode = ""
			ind++
		}
	}
	return list
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
	list := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		el, _ := strconv.Atoi(scanner.Text())
		list[i] = el
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

	tree.visit(tree.root)
	tree.code = tree.sb.String()

	fmt.Println("Got tree:", tree.sb.String())

	tree.generateCodeTable()
	fmt.Println("Got code table", tree.codeTable)

	tree.generateValueTable()
	fmt.Println("Got value table", tree.valueTable)

	listCode := tree.generateListCode(list)
	fmt.Println("Got list code:", listCode)

	fmt.Println("Got list:", tree.generateList(listCode, n))
}
