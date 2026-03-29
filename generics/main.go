package main

import "fmt"

type Number interface {
	int64 | float64
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	m := make(map[int64]int64)
	m[1] = 1
	m[2] = 2
	m[4] = 10
	fmt.Println(SumNumbers(m))
}
