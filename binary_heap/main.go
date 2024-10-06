package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Heap struct {
	list []int
}

func create(c int) *Heap {
	return &Heap{list: make([]int, 0, c)}
}

func (h *Heap) push(item int) {
	l := len(h.list)
	h.list = h.list[:l+1]
	h.list[l] = item
	h.sink()
}

func (h *Heap) sink() {
	l := len(h.list)
	i := l - 1
	if i == 0 {
		return
	}
	for {
		p := (i - 1) / 2
		if h.list[p] < h.list[i] {
			h.list[i], h.list[p] = h.list[p], h.list[i]
			i = p
		} else {
			return
		}

		if i == 0 {
			return
		}
	}
}

func (h *Heap) pop() int {
	l := len(h.list)
	last := l - 1
	el := h.list[last]
	h.list[0], h.list[last] = h.list[last], h.list[0]
	h.list = h.list[:last]
	h.dive()
	return el
}

func (h *Heap) dive() {
	i := 0
	cnt := len(h.list)

	if cnt == 0 {
		return
	}

	for {
		el := h.list[i]
		c := i

		l := 2*i + 1
		if l < cnt && el < h.list[l] {
			el = h.list[l]
			c = l
		}

		r := 2*i + 2
		if r < cnt && el < h.list[r] {
			el = h.list[r]
			c = r
		}

		if c == i {
			break
		}

		h.list[i], h.list[c] = h.list[c], h.list[i]
		i = c
	}
}

func (h *Heap) getSorted() []int {
	l := cap(h.list)
	return h.list[:l]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	cnt, _ := strconv.Atoi(scanner.Text())

	heap := create(cnt)

	for i := 0; i < cnt; i++ {
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		heap.push(v)
	}

	for len(heap.list) != 0 {
		heap.pop()
	}

	fmt.Println(heap.getSorted())
}
