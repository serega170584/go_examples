package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2, 4, 4, 4, 4, 4}
	k := 2

	res := make([]int, 0, len(nums))

	counts := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := counts[nums[i]]; !ok {
			res = append(res, nums[i])
		}
		counts[nums[i]]++
	}

	sort.Slice(res, func(i int, j int) bool {
		return counts[res[i]] <= counts[res[j]]
	})

	fmt.Println(res[len(res)-k:])
}
