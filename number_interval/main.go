package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4, 5, 6, 10, 10, 10, 10}
	fmt.Println(rbinarySearch(10, 10, list))
}

func binarySearch(x int, n int, list []int) int {
	l := 0
	r := n - 1
	for l < r {
		m := (l + r) / 2
		if check(m, x, list) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func rbinarySearch(x int, n int, list []int) int {
	l := 0
	r := n - 1
	for l < r {
		m := (l + r + 1) / 2
		if lcheck(m, x, list) {
			l = m + 1
		} else {
			l = r
		}
	}
	return l
}

func check(m int, x int, list []int) bool {
	return list[m] >= x
}

func lcheck(m int, x int, list []int) bool {
	return list[m] <= x
}
