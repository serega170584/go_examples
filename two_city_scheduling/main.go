package main

import (
	"fmt"
	"sort"
)

func main() {
	a := [][2]int{{1, 3}, {4, 8}, {9, 20}, {15, 60}, {21, 33}}
	second := 2
	d := make([]int, len(a))
	s := 0
	indexes := make([]int, len(a))
	for i, v := range a {
		d[i] = v[1] - v[0]
		s += v[0]
		indexes[i] = i
	}

	sort.Slice(indexes, func(i, j int) bool {
		return d[i] < d[j]
	})

	for i := len(d) - 1; i > len(d)-1-second; i-- {
		s += d[i]
	}

	fmt.Println(s)
}
