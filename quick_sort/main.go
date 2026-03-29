package main

import "fmt"

func main() {
	a := []int{5, 4, 3, 2, 1, 4, 6, 8, 1, 1, 1, 1}
	sort(a, 0, len(a)-1)
	fmt.Println(a)
}

func sort(a []int, low, high int) {
	if low >= high {
		return
	}

	i := partition(a, low, high)
	sort(a, low, i-1)
	sort(a, i+1, high)
}

func partition(a []int, low, high int) int {
	i := low - 1
	pivot := a[high]
	for j := low; j < high; j++ {
		if a[j] > pivot {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}

	i++
	a[i], a[high] = a[high], a[i]

	return i
}
