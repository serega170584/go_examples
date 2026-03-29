package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(getLargestList([]int{1, 3, 2, 4, 5, 7, 6, 9, 8, 3, 3, 3, 4, 4, 4, 4, 6, 7}, 3))
}

func getLargestList(arr []int, k int) []int {
	cntList := make(map[int]int)
	unique := make([]int, 0)
	for _, v := range arr {
		if _, ok := cntList[v]; !ok {
			unique = append(unique, v)
		}
		cntList[v]++
	}

	if len(unique) < k {
		return nil
	}

	sort.Slice(unique, func(i, j int) bool {
		cond := cntList[unique[i]] >= cntList[unique[j]]
		if cond && cntList[unique[i]] == cntList[unique[j]] {
			cond = cond && unique[i] < unique[j]
		}

		return cond
	})

	return unique[:k]
}
