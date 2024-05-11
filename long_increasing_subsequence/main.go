package main

import "fmt"

func main() {
	fmt.Println(maxDecIncSubsequence([]int{2, 4, 3, 2, 2}))
}

func maxDecIncSubsequence(a []int) int {
	if len(a) == 0 {
		return 0
	}

	if len(a) == 1 {
		return 1
	}

	maxIncDec := 1
	l := len(a)

	inc := 1
	dec := 1
	for i := 1; i < l; i++ {
		if a[i-1] < a[i] {
			inc++
		} else {
			inc = 1
		}

		if a[i-1] > a[i] {
			dec++
		} else {
			dec = 1
		}

		maxIncDec = max(maxIncDec, inc)
		maxIncDec = max(maxIncDec, dec)
	}

	maxIncDec = max(maxIncDec, inc)
	maxIncDec = max(maxIncDec, dec)

	return maxIncDec
}
