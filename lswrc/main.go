package main

import (
	"fmt"
)

// 1 3 4 7 9 11
// 2 5 6 8

// 2 5 | 6 8
// 1 3 4 | 7 9 11

func main() {
	a := []rune("abcaaabbcccdefg")
	m := make(map[rune]int)
	left := 0
	right := 0
	if len(a) == 0 {
		fmt.Println(0)
	}

	ml := 1
	m[a[0]] = 0

	right = 1

	for right < len(a) {
		if m[a[right]] >= left {
			left = m[a[right]] + 1
		}

		l := right - left + 1
		if l > ml {
			ml = l
		}

		m[a[right]] = right
		right++
	}

	fmt.Println(ml)
}
