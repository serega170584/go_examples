package main

import "fmt"

func main() {
	fmt.Println(lis([]int{1, 2, 5, 3, 4}))
}

func lis(a []int) int {
	if len(a) <= 1 {
		return 1
	}

	prev := a[0]
	maxLen := 1
	incLen := 1
	decLen := 1
	al := len(a)
	for i := 1; i < al; i++ {
		maxLen = max(maxLen, decLen)
		maxLen = max(maxLen, incLen)

		if a[i] > prev {
			incLen++
		} else {
			incLen = 1
		}

		if a[i] < prev {
			decLen++
		} else {
			decLen = 1
		}

		prev = a[i]
	}

	maxLen = max(maxLen, decLen)
	maxLen = max(maxLen, incLen)

	return maxLen
}
