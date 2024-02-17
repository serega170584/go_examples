package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Heap struct {
	capacity  int
	length    int
	list      []*Node
	positions map[int]int
}

type Node struct {
	visitor int
	time    int
	index   int
}

func NewHeap(capacity int) *Heap {
	list := make([]*Node, capacity)
	positions := make(map[int]int, capacity)
	return &Heap{capacity: capacity, list: list, positions: positions}
}

func (heap *Heap) push(node *Node) {
	i := heap.length
	heap.length++
	heap.list[i] = node
	heap.positions[node.visitor] = i
	heap.siftUp(i)
}

func (heap *Heap) findVisitorPosition(visitor int) int {
	if position, ok := heap.positions[visitor]; ok {
		return position
	}
	return -1
}

func (heap *Heap) change(i int, node *Node) {
	heap.list[i] = node
	heap.heapify(i)
	heap.siftUp(i)
}

func (heap *Heap) find(node *Node) int {
	if position, ok := heap.positions[node.visitor]; ok {
		return position
	}
	return -1
}

func (heap *Heap) heapify(i int) {
	current := i
	lastIndex := heap.length - 1
	oldCurrent := lastIndex
	for current != oldCurrent {
		oldCurrent = current
		left := 2*current + 1
		minIndex := current
		if left <= lastIndex && heap.list[minIndex].time > heap.list[left].time {
			minIndex = left
		}

		right := 2*current + 2
		if right <= lastIndex && heap.list[minIndex].time > heap.list[right].time {
			minIndex = right
		}

		heap.list[minIndex], heap.list[current] = heap.list[current], heap.list[minIndex]
		heap.positions[heap.list[minIndex].visitor] = minIndex
		heap.positions[heap.list[current].visitor] = current

		if current != minIndex {
			current = minIndex
		}
	}
}

func (heap *Heap) siftUp(i int) {
	for i != 0 {
		parent := (i - 1) / 2
		if heap.list[parent].time > heap.list[i].time {
			heap.list[parent], heap.list[i] = heap.list[i], heap.list[parent]
			heap.positions[heap.list[parent].visitor] = parent
			heap.positions[heap.list[i].visitor] = i
		}
		i = parent
	}
}

func (heap *Heap) pop() *Node {
	node := heap.list[0]
	i := heap.length - 1
	heap.length--
	heap.list[0] = heap.list[i]
	delete(heap.positions, node.visitor)
	heap.heapify(0)
	return node
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	fmt.Println("Enter events count")
	scanner.Scan()
	eventsCount, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter list count")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter list")
	list := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		list[i], _ = strconv.Atoi(scanner.Text())
	}

	heap := NewHeap(eventsCount)
	cache := make(map[int]*Node, eventsCount)
	for i := 0; i < eventsCount; i++ {
		node := &Node{visitor: list[i], time: i, index: i}
		cache[list[i]] = node
		heap.push(node)
	}

	for i := eventsCount; i < n; i++ {
		newNode := &Node{visitor: list[i], time: i}
		index := heap.find(newNode)
		if index == -1 {
			node := heap.pop()
			heap.push(newNode)
			delete(cache, node.visitor)
		} else {
			heap.change(index, newNode)
		}
		cache[list[i]] = newNode
		for _, v := range cache {
			fmt.Print(v.visitor)
		}
		fmt.Println(cache)
	}
}
