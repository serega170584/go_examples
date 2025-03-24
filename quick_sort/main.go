package main

import "fmt"

func main() {
	a := []int{5, 4, 3, 2, 1}
	sort(a, 0, len(a)-1)
	fmt.Println(a)
}

func sort(a []int, low int, high int) {
	if low >= high {
		return
	}

	i := partition(a, low, high)
	sort(a, low, i-1)
	sort(a, i+1, high)
}

func partition(a []int, low int, high int) int {
	pivot := a[high]
	i := low - 1
	for j := low; j < high; j++ {
		if a[j] < pivot {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}

	i++
	a[high], a[i] = a[i], a[high]

	return i
}
