package main

import "fmt"

func main() {
	fmt.Println(union([]int{1, 2, 3, 7, 6, 4}, []int{6, 4, 3, 2, 90}, []int{7, 9, 3, 2, 1, 89, 90}, []int{3, 4, 5, 90}))
}

func union(a ...[]int) [][]int {
	l := len(a[0])
	al := len(a)
	for _, v := range a {
		if len(v) < l {
			l = len(v)
		}
	}

	u := make([][]int, 0, l)
	for i := 0; i < l; i++ {
		el := make([]int, 0, al)
		for j := 0; j < al; j++ {
			el = append(el, a[j][i])
		}
		u = append(u, el)
	}

	return u
}
