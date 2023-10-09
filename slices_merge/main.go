package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// 6 5 4 3 2 1
// _ 5 4 3 2 1
//     _
//  5      4
// 3  2   1
// 5 4 3 2 1

type Node struct {
	val  int
	next *Node
}

type BinaryHeap struct {
	list []*Node
	cnt  int
}

func newBinaryHeap(capacity int) *BinaryHeap {
	heap := &BinaryHeap{}
	heap.list = make([]*Node, capacity)
	return heap
}

func (heap *BinaryHeap) add(node *Node) {
	cnt := heap.cnt

	heap.list[cnt] = node

	heap.heapify(cnt)
}

func (heap *BinaryHeap) incCnt() {
	heap.cnt++
}

func (heap *BinaryHeap) heapify(curPointer int) {
	for curPointer != 0 {
		parentPointer := (curPointer+1)/2 - 1
		if heap.list[curPointer].val < heap.list[parentPointer].val {
			heap.list[curPointer], heap.list[parentPointer] = heap.list[parentPointer], heap.list[curPointer]
		}
		curPointer = parentPointer
	}
}

func (heap *BinaryHeap) build() {
	for i := heap.cnt - 1; i > 0; i-- {
		heap.heapify(i)
	}
}

func (heap *BinaryHeap) pop() (*Node, bool) {
	if heap.cnt == 0 {
		return nil, false
	}
	return heap.list[0], true
}

func (heap *BinaryHeap) recalcCnt() {
	heap.list[0] = heap.list[heap.cnt-1]
	heap.cnt--
}

func (heap *BinaryHeap) push(node *Node) {
	heap.list[0] = node
	pointer := 0
	for pointer < heap.cnt {
		left := pointer*2 + 1
		right := pointer*2 + 2

		leftVal := int(math.Inf(1))
		if left < heap.cnt {
			leftVal = heap.list[left].val
		}

		rightVal := int(math.Inf(1))
		if right < heap.cnt {
			rightVal = heap.list[right].val
		}

		if leftVal >= node.val && node.val <= rightVal {
			break
		}

		if leftVal < rightVal {
			heap.list[pointer], heap.list[left] = heap.list[left], heap.list[pointer]
			pointer = left
		} else {
			heap.list[pointer], heap.list[right] = heap.list[right], heap.list[pointer]
			pointer = right
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	cntStr := scanner.Text()
	cnt, _ := strconv.Atoi(cntStr)

	nodes := make([]*Node, cnt)
	for i := 0; i < cnt; i++ {
		scanner.Scan()
		curCnt, _ := strconv.Atoi(scanner.Text())

		var prev *Node

		for j := 0; j < curCnt; j++ {
			scanner.Scan()

			node := &Node{}
			node.val, _ = strconv.Atoi(scanner.Text())

			if prev == nil {
				nodes[i] = node
			} else {
				prev.next = node
			}

			prev = node
		}
	}

	heap := newBinaryHeap(cnt)

	for _, node := range nodes {
		heap.add(node)
		heap.incCnt()
	}

	fmt.Println(heap)

	for _, node := range heap.list {
		fmt.Println(node)
	}

	for node, ok := heap.pop(); ok; node, ok = heap.pop() {
		fmt.Println(node.val)
		node1 := node.next
		if node1 == nil {
			heap.recalcCnt()
			heap.build()
		} else {
			heap.push(node1)
		}
	}
}
