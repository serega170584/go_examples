package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type BinaryHeap struct {
	list []int
	cnt  int
}

func newBinaryHeap(capacity int) *BinaryHeap {
	heap := &BinaryHeap{}
	heap.list = make([]int, capacity)
	return heap
}

func (heap *BinaryHeap) add(el int) {
	heap.list[heap.cnt] = el
	heap.heapify(heap.cnt)
	heap.cnt++
}

func (heap *BinaryHeap) heapify(ind int) {
	for ind != 0 {
		parentInd := (ind+1)/2 - 1
		if heap.list[parentInd] > heap.list[ind] {
			heap.list[parentInd], heap.list[ind] = heap.list[ind], heap.list[parentInd]
		}
		ind = parentInd
	}
}

func (heap *BinaryHeap) min() int {
	return heap.list[0]
}

func (heap *BinaryHeap) push(val int) {
	pointer := 0
	heap.list[pointer] = val
	for true {
		left := 2*pointer + 1
		right := 2*pointer + 2

		leftVal := int(math.Inf(1))
		if left < heap.cnt {
			leftVal = heap.list[left]
		}

		rightVal := int(math.Inf(1))
		if right < heap.cnt {
			rightVal = heap.list[right]
		}

		if leftVal >= heap.list[pointer] && heap.list[pointer] <= rightVal {
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
	cnt, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	heap := newBinaryHeap(k)

	for i := 0; i < k; i++ {
		scanner.Scan()
		val, _ := strconv.Atoi(scanner.Text())

		heap.add(val)
	}

	min := heap.min()

	for i := k; i < cnt; i++ {
		scanner.Scan()
		val, _ := strconv.Atoi(scanner.Text())

		if val > min {
			heap.push(val)
		}

		min = heap.min()
	}

	fmt.Println(min)
}
