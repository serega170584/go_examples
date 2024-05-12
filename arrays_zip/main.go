package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(arrayZip([]int{1, 2, 2}, []int{2, 3}))
}

func arrayZip(list ...[]int) [][]int {
	l := math.MaxInt
	for _, a := range list {
		l = min(l, len(a))
	}

	res := make([][]int, l)

	ll := len(list)
	for i := 0; i < l; i++ {
		res[i] = make([]int, ll)
	}

	for i := 0; i < l; i++ {
		for j := 0; j < ll; j++ {
			res[i][j] = list[j][i]
		}
	}

	return res
}
