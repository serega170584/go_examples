package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(mergeSlices([]int{1, 2, 3, 4}, []int{5, 6, 7}, []int{8, 9}))
}

func mergeSlices(list ...[]int) [][]int {
	minLength := math.MaxInt

	for _, s := range list {
		minLength = min(minLength, len(s))
	}

	res := make([][]int, minLength)
	for i := 0; i < minLength; i++ {
		res[i] = make([]int, len(list))
		for j := 0; j < len(list); j++ {
			res[i][j] = list[j][i]
		}
	}

	return res
}
