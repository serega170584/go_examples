package main

import "fmt"

func main() {
	s := []rune{'a', 'b', 'c', 'a', 'b', 'c', 'b', 'b'}

	if len(s) <= 1 {
		fmt.Println(len(s))
	}

	left := 0
	l := 1
	m := make(map[rune]int)
	m[s[0]] = 0
	for right := 1; right < len(s); right++ {
		v := s[right]
		if i, ok := m[v]; ok && i >= left {
			left = i + 1
		}

		l = max(l, right-left+1)

		m[v] = right
	}

	fmt.Println(l)
}
