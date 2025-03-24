package main

import "fmt"

func merge(a []int, b []int) []int {
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}
	res := make([]int, 0, len(a)+len(b))
	i, j := 0, 0

	for i < len(a) || j < len(b) {
		if i == len(a) || (j < len(b) && a[i] > b[j]) {
			res = append(res, b[j])
			j++
		} else {
			res = append(res, a[i])
			i++
		}
	}
	return res
}

func main() {
	b := []int{1, 3, 3}
	a := []int{3, 4, 5, 6, 7, 8}

	fmt.Println(merge(a, b))
}
