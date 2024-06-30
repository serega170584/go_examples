package main

import "fmt"

func main() {
	a1 := []int{1, 2, 3, 4, 8, 9, 10, 13, 15, 16}
	a2 := []int{3, 5, 6, 11, 12, 15, 18, 20}
	fmt.Println(diff(a1, a2))
}

func diff(a1 []int, a2 []int) []int {
	p := 0
	d := make([]int, 0, len(a1))
	for _, v := range a1 {
		for p != len(a2) && a2[p] < v {
			p++
		}

		if p == len(a2) || a2[p] > v {
			d = append(d, v)
		}
	}

	return d
}
