package main

import "fmt"

func main() {
	fmt.Println(mergeSlices([]int{3, 3, 4, 5}, []int{1, 2, 6}))
}

func mergeSlices(a []int, b []int) []int {
	al := len(a)
	bl := len(b)
	c := make([]int, 0, al+bl)
	p := 0
	for _, v := range a {
		for p != bl && b[p] < v {
			c = append(c, b[p])
			p++
		}

		if p == bl || v <= b[p] {
			c = append(c, v)
		}
	}

	return c
}
