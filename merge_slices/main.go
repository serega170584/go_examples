package main

import "fmt"

func main() {
	fmt.Println(mergeSlices([]int{1, 2, 6}, []int{3, 4, 5}))
}

func mergeSlices(a []int, b []int) []int {
	al := len(a)
	bl := len(b)
	var res []int
	if al < bl {
		for i := 0; i < al; i++ {
			b = append(b, 0)
		}
		res = b
	} else {
		for i := 0; i < bl; i++ {
			a = append(a, 0)
		}
		res = a
	}

	l := len(res)

	ai := al - 1
	bi := bl - 1
	for i := l - 1; i >= 0; i-- {
		if ai == -1 {
			res[i] = b[bi]
			bi--
			continue
		}

		if bi == -1 {
			res[i] = a[ai]
			ai--
			continue
		}

		if a[ai] < b[bi] {
			res[i] = b[bi]
			bi--
		} else {
			res[i] = a[ai]
			ai--
		}
	}
	return res
}
