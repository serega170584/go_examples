package main

import "fmt"

func main() {
	s := []rune{'a', 'b', 'c', 'a', 'b', 'c', 'b', 'b'}

	if len(s) == 0 {
		fmt.Println("")
	}

	left := 0
	res := 1
	m := make(map[rune]int)
	m[s[left]] = left
	for right := 1; right < len(s); right++ {
		if i, ok := m[s[right]]; ok && i >= left {
			left = i + 1
		}

		v := right - left + 1
		res = max(res, v)

		m[s[right]] = right
	}

	fmt.Println(res)
}
