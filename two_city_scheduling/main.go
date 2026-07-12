package main

import (
	"fmt"
	"sort"
)

func main() {
	a := [][2]int{{1, 3}, {4, 8}, {9, 20}, {15, 60}, {21, 33}}
	second := 2
	sum := 0
	d := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		sum += a[i][0]
		d[i] = i
	}

	sort.Slice(d, func(i, j int) bool {
		left := d[i]
		right := d[j]
		return a[left][0]-a[left][1] < a[right][0]-a[right][1]
	})

	for i := 0; i < second; i++ {
		sum -= a[d[i]][0] - a[d[i]][1]
	}

	fmt.Println(sum)
}
