package main

import "fmt"

func main() {
	a := []int{5, 4, 3, 2, 2, 2, 2, 1}
	sort(a, 0, len(a)-1)
	fmt.Println(a)
}

func sort(a []int, lo int, hi int) {
	if lo >= hi {
		return
	}

	j := partition(a, lo, hi)
	sort(a, lo, j-1)
	sort(a, j+1, hi)
}

func partition(a []int, lo int, hi int) int {
	v := a[lo]
	i := lo
	j := hi + 1
	for {
		i++
		for a[i] < v {
			if i == hi {
				break
			}
			i++
		}

		j--
		for a[j] > v {
			if j == lo {
				break
			}
			j--
		}

		if i >= j {
			break
		}

		a[i], a[j] = a[j], a[i]
	}

	a[lo], a[j] = a[j], a[lo]

	return j
}
