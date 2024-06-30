package main

import "fmt"

func main() {
	fmt.Println(lis([]int{1, 2, 5, 3, 4}))
}

func lis(a []int) int {
	if len(a) == 0 {
		return 0
	}

	if len(a) == 1 {
		return 1
	}

	inc := 1
	dec := 1
	m := 1
	for i := 1; i < len(a); i++ {
		if a[i] > a[i-1] {
			inc++
			m = max(inc, m)
		} else {
			inc = 1
		}

		if a[i] < a[i-1] {
			dec++
			m = max(dec, m)
		} else {
			dec = 1
		}
	}

	return m
}
