package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	val  int
	next *Node
}

type Heap struct {
	cnt int
	arr []*Node
}

func (h *Heap) add(el *Node) {
	if len(h.arr) == h.cnt {
		h.arr = append(h.arr, el)
	} else {
		h.arr[h.cnt] = el
	}

	h.cnt++

	i := h.cnt - 1
	parent := (i - 1) / 2
	for i > 0 && h.arr[parent].val > h.arr[i].val {
		h.arr[parent], h.arr[i] = h.arr[i], h.arr[parent]
		i = parent
		parent = (i - 1) / 2
	}
}

func (h *Heap) build(arr []*Node) {
	for _, val := range arr {
		h.add(val)
	}
}

func (h *Heap) getMin() *Node {
	min := h.arr[0]
	h.arr[0] = h.arr[h.cnt-1]
	h.cnt--
	h.heapify(0)
	return min
}

func (h *Heap) isEmpty() bool {
	return h.cnt == 0
}

func (h *Heap) heapify(i int) {
	var left, right, smallest int
	for {
		left = 2*i + 1
		right = 2*i + 2

		if left < h.cnt && h.arr[i].val > h.arr[left].val {
			smallest = left
		}

		if right < h.cnt && h.arr[smallest].val > h.arr[right].val {
			smallest = right
		}

		if smallest == i {
			break
		}

		h.arr[i], h.arr[smallest] = h.arr[smallest], h.arr[i]
		i = smallest
	}
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	s.Scan()

	cnt, _ := strconv.Atoi(s.Text())
	linkedLists := make([]*Node, cnt)
	totalListCnt := 0
	for i := 0; i < cnt; i++ {
		s.Scan()
		listCnt, _ := strconv.Atoi(s.Text())
		totalListCnt += listCnt
		var prev *Node
		for j := 0; j < listCnt; j++ {
			s.Scan()
			node := &Node{}
			node.val, _ = strconv.Atoi(s.Text())

			if j == 0 {
				linkedLists[i] = node
			}

			if prev != nil {
				prev.next = node
			}

			prev = node
		}
	}

	h := &Heap{}
	for _, node := range linkedLists {
		h.add(node)
	}

	res := make([]int, totalListCnt)
	pointer := 0

	for !h.isEmpty() {
		min := h.getMin()
		res[pointer] = min.val
		pointer++

		if min.next != nil {
			h.add(min.next)
		}
	}

	fmt.Println(res)
}
