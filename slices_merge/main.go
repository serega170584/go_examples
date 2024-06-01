package main

import "fmt"

func main() {
	fmt.Println(mergeSlices([]int{1, 2, 3}, []int{3, 2, 1, 4}, []int{2, 1}, []int{1, 1}))
}

func mergeSlices(list ...[]int) [][]int {
	l := len(list[0])
	ll := len(list)
	for i := 1; i < ll; i++ {
		l = min(l, len(list[i]))
	}

	res := make([][]int, l)
	for i := 0; i < l; i++ {
		for j := 0; j < ll; j++ {
			res[i] = append(res[i], list[j][i])
		}
	}

	return res
}
