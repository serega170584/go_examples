package main

import "fmt"

func main() {
	A := 239
	fmt.Println(colorful(A))
}

func colorful(A int) int {
	parts := make([]int, 0)
	m := A % 10
	A /= 10
	parts = append(parts, m)
	for A != 0 {
		m = A % 10
		A /= 10
		parts = append([]int{m}, parts...)
	}

	l := len(parts)
	seqMap := make(map[int]int)
	for i := 0; i < l; i++ {
		num := parts[i]
		offset := 0
		seqMap[num]++
		if seqMap[num] > 1 {
			return 0
		}
		for offset = 1; offset < l-i; offset++ {
			num = num * parts[i+offset]
			seqMap[num]++
			if seqMap[num] > 1 {
				return 0
			}
		}
	}

	return 1
}
