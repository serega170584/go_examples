package main

import "fmt"

func main() {
	fmt.Println(zip([]int{1, 2}, []int{3, 4, 5}, []int{5, 6, 7, 8, 9}))
}

func zip(a ...[]int) [][]int {
	l := len(a)
	ml := len(a[0])
	for i := 1; i < l; i++ {
		ml = min(ml, len(a[i]))
	}

	res := make([][]int, ml)

	for i := 0; i < ml; i++ {
		res[i] = make([]int, 0)
		for j := 0; j < l; j++ {
			res[i] = append(res[i], a[j][i])
		}
	}

	return res
}
