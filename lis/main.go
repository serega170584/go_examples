package main

import "fmt"

func main() {
	fmt.Println(lis([]int{1, 6, 5, 4, 3}))
}

func lis(a []int) int {
	if len(a) == 1 {
		return 1
	}

	l := len(a)
	m := 1
	inc := 1
	dec := 1
	for i := 1; i < l; i++ {
		if a[i] > a[i-1] {
			inc++
		} else {
			inc = 1
		}

		if a[i] < a[i-1] {
			dec++
		} else {
			dec = 1
		}

		m = max(inc, m)
		m = max(dec, m)
	}

	return m
}
