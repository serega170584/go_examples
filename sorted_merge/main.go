package main

import "fmt"

func main() {
	fmt.Println(mergeSorted([]int{1, 3, 5, 6}, []int{2, 4, 6, 8, 9, 10}))
}

func mergeSorted(a []int, b []int) []int {
	al := len(a)
	bl := len(b)
	l := al + bl
	c := make([]int, 0, l)

	ai := 0
	bi := 0
	for i := 0; i < l; i++ {
		if ai == al {
			c = append(c, b[bi])
			bi++
			continue
		}

		if bi == bl {
			c = append(c, a[i])
			ai++
			continue
		}

		if a[ai] < b[bi] {
			c = append(c, a[ai])
			ai++
		} else {
			c = append(c, b[bi])
			bi++
		}
	}

	return c
}
