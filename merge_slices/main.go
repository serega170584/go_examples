package main

import "fmt"

func main() {
	fmt.Println(mergeSlices([]int{3, 3, 4, 5}, []int{1, 2, 6}))
}

func mergeSlices(a []int, b []int) []int {
	i := 0
	j := 0
	res := make([]int, 0, len(a)+len(b))
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			res = append(res, a[i])
			i++
			continue
		}

		if a[i] > b[j] {
			res = append(res, b[j])
			j++
			continue
		}

		res = append(res, a[i], b[j])
		i++
		j++
	}

	if i < len(a) {
		res = append(res, a[i:]...)
	}

	if j < len(b) {
		res = append(res, b[j:]...)
	}

	return res
}
