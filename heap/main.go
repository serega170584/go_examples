package main

import "math"

type heap struct {
	list []int
	cnt  int
}

func (h *heap) add(e int) {
	h.list[h.cnt] = e
	h.flow()
	h.cnt++
}

func (h *heap) flow() {
	i := h.cnt
	for i != 0 {
		pi := (i - 1) / 2
		if h.list[i] <= h.list[pi] {
			return
		} else {
			h.list[i], h.list[pi] = h.list[pi], h.list[i]
		}
		i = pi
	}
}

func (h *heap) pop() int {
	if h.cnt == 0 {
		return -1
	}

	e := h.list[0]
	if h.cnt != 1 {
		h.list[0] = h.list[h.cnt-1]
	}
	h.cnt--

	h.dive()

	return e
}

func (h *heap) dive() {
	i := 0
	for i < h.cnt {
		maxI := i

		left := math.MinInt
		li := 2*i + 1
		if li < h.cnt {
			left = h.list[li]
		}
		if left > h.list[i] {
			maxI = li
		}

		right := math.MinInt
		ri := 2*i + 2
		if ri < h.cnt {
			right = h.list[ri]
		}
		if right > h.list[maxI] {
			maxI = ri
		}

		h.list[maxI], h.list[i] = h.list[i], h.list[maxI]

		if i == maxI {
			break
		}

		i = maxI
	}
}

func main() {

}
