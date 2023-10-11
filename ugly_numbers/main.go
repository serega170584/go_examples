package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Node struct {
	val    int
	factor int
	next   *Node
}

type BinaryHeap struct {
	val  int
	list []*Node
	cnt  int
}

func newHeap(capacity int) *BinaryHeap {
	heap := &BinaryHeap{}
	heap.list = make([]*Node, capacity)
	return heap
}

func (heap *BinaryHeap) add(node *Node) {
	heap.list[heap.cnt] = node
	heap.heapify(heap.cnt)
	heap.cnt++
}

func (heap *BinaryHeap) push(node *Node) {
	heap.list[0] = node
	ind := 0
	for true {
		left := 2*ind + 1
		right := 2*ind + 2

		leftVal := int(math.Inf(1))
		if left < heap.cnt {
			leftVal = heap.list[left].val
		}

		rightVal := int(math.Inf(1))
		if right < heap.cnt {
			rightVal = heap.list[right].val
		}

		if leftVal >= heap.list[ind].val && heap.list[ind].val <= rightVal {
			break
		}

		if leftVal < rightVal {
			heap.list[ind], heap.list[left] = heap.list[left], heap.list[ind]
		} else {
			heap.list[ind], heap.list[right] = heap.list[right], heap.list[ind]
		}
	}
}

func (heap *BinaryHeap) heapify(ind int) {
	for ind != 0 {
		parent := (ind+1)/2 - 1
		if heap.list[ind].val < heap.list[parent].val {
			heap.list[ind], heap.list[parent] = heap.list[parent], heap.list[ind]
		}
		ind = parent
	}
}

func (heap *BinaryHeap) getMin() *Node {
	return heap.list[0]
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	var cnt int
	var min *Node

	uglyArr := [3]int{2, 3, 5}

	twoFactorHeap := newHeap(3)
	threeFactorHeap := newHeap(3)
	fiveFactorHeap := newHeap(3)

	heap := newHeap(3)

	for _, val := range uglyArr {
		twoFactorNode := &Node{val: val, factor: 2}
		if val == 2 {
			heap.add(twoFactorNode)
		} else {
			twoFactorHeap.add(twoFactorNode)
		}

		threeFactorNode := &Node{val: val, factor: 3}
		if val == 2 {
			heap.add(threeFactorNode)
		} else {
			threeFactorHeap.add(threeFactorNode)
		}

		fiveFactorNode := &Node{val: val, factor: 5}
		if val == 2 {
			heap.add(fiveFactorNode)
		} else {
			fiveFactorHeap.add(fiveFactorNode)
		}
	}

	for cnt < n {
		curMin := heap.getMin()
		if min == nil || min.val != curMin.val {
			min = curMin
			cnt++
		}

		if curMin.factor == 2 {
			node := &Node{val: curMin.val * 2, factor: 2}
			twoFactorMin := twoFactorHeap.getMin()
			twoFactorHeap.push(node)
			heap.push(twoFactorMin)
		}

		if curMin.factor == 3 {
			node := &Node{val: curMin.val * 3, factor: 3}
			threeFactorMin := threeFactorHeap.getMin()
			threeFactorHeap.push(node)
			heap.push(threeFactorMin)
		}

		if curMin.factor == 5 {
			node := &Node{val: curMin.val * 5, factor: 5}
			fiveFactorMin := fiveFactorHeap.getMin()
			fiveFactorHeap.push(node)
			heap.push(fiveFactorMin)
		}
	}

	fmt.Println(min.val)

}
