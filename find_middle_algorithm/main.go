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

	fmt.Println("Enter list count")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter window size")
	scanner.Scan()
	windowSize, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter list")
	list := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		list[i], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println("Got middle", findMiddle(n, list, windowSize))
}

type Heap struct {
	length      int
	capacity    int
	list        []*Node
	positionMap map[int]*Node
	isMaxTop    bool
}

type Node struct {
	val      int
	index    int
	position int
}

func NewHeap(capacity int, windowSize int, isMaxTop bool) *Heap {
	list := make([]*Node, capacity)
	positionMap := make(map[int]*Node, windowSize)
	return &Heap{capacity: capacity, list: list, positionMap: positionMap, isMaxTop: isMaxTop}
}

func (heap *Heap) push(node *Node) {
	i := heap.length
	heap.length++
	heap.list[i] = node
	heap.positionMap[node.position] = node
	node.index = i
	heap.siftUp(i)
}

func (heap *Heap) pop() *Node {
	i := heap.length - 1
	el := heap.list[0]
	heap.list[0] = heap.list[i]
	heap.list[0].index = 0
	delete(heap.positionMap, el.position)
	heap.length--
	heap.heapify(0)
	return el
}

func (heap *Heap) find(position int) int {
	if node, ok := heap.positionMap[position]; ok {
		delete(heap.positionMap, position)
		return node.index
	}

	return -1
}

func (heap *Heap) change(i int, node *Node) {
	heap.list[i] = node
	node.index = i
	heap.positionMap[node.position] = node
	heap.heapify(i)
	heap.siftUp(i)
}

func (heap *Heap) heapify(i int) {
	current := i
	last := heap.length - 1
	currentOld := heap.length - 1

	for current != currentOld {
		currentOld = current
		minInd := current

		left := 2*current + 1
		if heap.isMaxTop {
			if left <= last && heap.list[left].val > heap.list[minInd].val {
				minInd = left
			}
		} else {
			if left <= last && heap.list[left].val < heap.list[minInd].val {
				minInd = left
			}
		}

		right := 2*current + 2
		if heap.isMaxTop {
			if right <= last && heap.list[right].val > heap.list[minInd].val {
				minInd = right
			}
		} else {
			if right <= last && heap.list[right].val < heap.list[minInd].val {
				minInd = right
			}
		}

		heap.list[minInd], heap.list[current] = heap.list[current], heap.list[minInd]
		heap.list[minInd].index = minInd
		heap.list[current].index = current
		current = minInd
	}
}

func (heap *Heap) siftUp(i int) {
	for i != 0 {
		parent := (i - 1) / 2
		if heap.isMaxTop {
			if heap.list[i].val > heap.list[parent].val {
				heap.list[i], heap.list[parent] = heap.list[parent], heap.list[i]
			}
		} else {
			if heap.list[i].val < heap.list[parent].val {
				heap.list[i], heap.list[parent] = heap.list[parent], heap.list[i]
			}
		}
		heap.list[parent].index = parent
		heap.list[i].index = i
		i = parent
	}
}

func findMiddle(n int, list []int, windowSize int) []int {
	leastHeapSize := windowSize / 2
	mostHeapSize := windowSize/2 + windowSize%2
	leastHeap := NewHeap(leastHeapSize, windowSize, true)
	mostHeap := NewHeap(mostHeapSize, windowSize, false)
	middle := make([]int, n)
	middleIndex := 0
	for i, v := range list {
		node := &Node{val: v, position: i}
		if i >= windowSize {
			searchedHeap := leastHeap
			searched := leastHeap.find(i - windowSize)
			if searched == -1 {
				searchedHeap = mostHeap
				searched = mostHeap.find(i - windowSize)
			}
			searchedHeap.change(searched, node)
		} else {
			if i%2 == 0 {
				mostHeap.push(node)
			} else {
				leastHeap.push(node)
			}
		}

		if leastHeap.length != 0 {
			leastHeapNode := leastHeap.pop()
			mostHeapNode := mostHeap.pop()
			if leastHeapNode.val > mostHeapNode.val {
				leastHeapNode, mostHeapNode = mostHeapNode, leastHeapNode
			}
			leastHeap.push(leastHeapNode)
			mostHeap.push(mostHeapNode)
		}

		node = mostHeap.pop()
		middle[middleIndex] = node.val
		mostHeap.push(node)
		middleIndex++
	}
	return middle
}
