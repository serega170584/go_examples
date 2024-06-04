package main

import "fmt"

func main() {
	fmt.Println(mergeSlices([]int{3, 3, 4, 5}, []int{1, 2, 6}))
}

func mergeSlices(a []int, b []int) []int {
	al := len(a)
	bl := len(b)
	cl := al + bl
	c := make([]int, cl)

	ai := 0
	bi := 0
	for i := 0; i < cl; i++ {
		if ai == al {
			c[i] = b[bi]
			bi++
			continue
		}
		if bi == bl {
			c[i] = a[ai]
			ai++
			continue
		}
		if a[ai] > b[bi] {
			c[i] = b[bi]
			bi++
			continue
		}

		c[i] = a[ai]
		ai++
	}
	return c
}
