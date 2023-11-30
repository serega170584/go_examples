package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Heap struct {
	list []int
	cnt  int
}

func NewHeap(capacity int) *Heap {
	heap := &Heap{}
	heap.list = make([]int, 0, capacity)
	return heap
}

func (h *Heap) add(v int) {
	h.cnt++
	i := h.cnt - 1
	h.list = append(h.list, v)
	for i != 0 {
		p := (i - 1) / 2

		if h.list[p] > h.list[i] {
			h.list[i], h.list[p] = h.list[p], h.list[i]
			i = p
			continue
		}

		break
	}
}

func (h *Heap) getMin() int {
	min := -1
	if h.cnt > 0 {
		min = h.list[0]
		h.cnt--
	}
	if h.cnt > 0 {
		h.list[0] = h.list[h.cnt]
		h.heapify()
	}
	return min
}

func (h *Heap) isEmpty() bool {
	return h.cnt == 0
}

func (h *Heap) heapify() {
	i := 0
	for {
		smallest := i

		l := 2*i + 1
		if l > h.cnt-1 {
			break
		}

		if h.list[smallest] > h.list[l] {
			smallest = l
		}

		r := 2*i + 2
		if r < h.cnt && h.list[smallest] > h.list[r] {
			smallest = r
		}

		if smallest == i {
			break
		}

		h.list[smallest], h.list[i] = h.list[i], h.list[smallest]
		i = smallest
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()

	cnt, _ := strconv.Atoi(scanner.Text())
	heap := NewHeap(cnt)
	for i := 0; i < cnt; i++ {
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		heap.add(v)
	}

	res := make([]string, 0, cnt)
	for !heap.isEmpty() {
		v := strconv.Itoa(heap.getMin())
		res = append(res, v)
	}
	fmt.Println(strings.Join(res, " "))
}
