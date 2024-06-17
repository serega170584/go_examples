package main

import "fmt"

func main() {
	fmt.Println(diff([]int{3, 4, 5, 6, 7}, []int{1, 2, 6, 7}))
}

func diff(a []int, b []int) []int {
	i := 0
	al := len(a)
	bl := len(b)
	res := make([]int, 0, al)
	for _, v := range a {
		for i != bl && v > b[i] {
			i++
		}

		if i == bl || v != b[i] {
			res = append(res, v)
		}
	}

	return res
}
