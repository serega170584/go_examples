package main

import (
	"fmt"
	"sort"
)

func main() {
	n := 9
	k := 3
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := []int{10, 11, 12, 13, 200, 14, 15, 16, 17}

	diff := make([]int, n)
	for i := 0; i < n; i++ {
		diff[i] = i
	}

	sort.Slice(diff, func(i, j int) bool {
		return b[i]-a[i] < b[j]-a[j]
	})

	s := 0
	for i := 0; i < k; i++ {
		currIdx := diff[i]
		s += a[currIdx]
	}

	for i := k; i < n; i++ {
		currIdx := diff[i]
		s += b[currIdx]
	}

	fmt.Println(s)
}
