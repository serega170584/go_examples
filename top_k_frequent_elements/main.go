package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2, 4, 4, 4, 4, 4}
	k := 2

	m := make(map[int]int)
	items := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; !ok {
			items = append(items, nums[i])
		}
		m[nums[i]]++
	}

	sort.Slice(items, func(i, j int) bool {
		iv := items[i]
		jv := items[j]

		return m[iv] < m[jv]
	})

	fmt.Println(items[len(items)-k:])
}
