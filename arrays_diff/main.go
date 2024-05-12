package main

import "fmt"

func main() {
	fmt.Println(diff([]int{1, 4, 7, 10, 13}, []int{2, 3, 6, 7, 9}))
}

func diff(a []int, b []int) []int {
	bi := 0
	res := make([]int, 0, len(a))
	bl := len(b)
	for _, v := range a {
		for bi != bl && b[bi] < v {
			bi++
		}

		if bi == bl {
			res = append(res, v)
			continue
		}

		if v != b[bi] {
			res = append(res, v)
		}
	}

	return res
}
