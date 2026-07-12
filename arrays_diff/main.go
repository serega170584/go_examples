package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 4, 5, 6, 7, 10}
	b := []int{8, 8, 8, 9, 9, 10}

	i := 0
	j := 0
	res := make([]int, 0)
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			res = append(res, a[i])
			i++
			continue
		}

		if a[i] > b[j] {
			j++
			continue
		}

		if a[i] == b[j] {
			i++
			j++
		}
	}

	fmt.Println(res)
}
