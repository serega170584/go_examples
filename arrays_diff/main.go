package main

import "fmt"

func main() {
	fmt.Println(diff([]int{1, 2, 3, 4, 5, 10, 11}, []int{4, 5, 6, 7, 8, 9, 10}))
}

func diff(a []int, b []int) []int {
	bp := 0
	al := len(a)
	bl := len(b)
	res := make([]int, 0, al)
	for _, v := range a {
		for bp != bl && v > b[bp] {
			bp++
		}

		if bp == bl || v < b[bp] {
			res = append(res, v)
		}
	}

	return res
}
