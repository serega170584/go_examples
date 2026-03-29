package main

import (
	"fmt"
)

func main() {
	a := []int{23, 3, 1, 2}
	b := []int{6, 2, 4, 23}

	// [2, 23]
	fmt.Printf("%v\n", intersection(a, b))
	a = []int{1, 1, 1}
	b = []int{1, 1, 1, 1}

	// [1, 1, 1]
	fmt.Printf("%v\n", intersection(a, b))
}

func intersection(a, b []int) []int {
	var res []int
	counter := make(map[int]int)
	for _, val := range a {
		counter[val]++
	}

	for _, val := range b {
		if count, ok := counter[val]; ok && counter > 0 {
			res = append(res, val)
			counter[val]--
		}
	}

	return res

}
