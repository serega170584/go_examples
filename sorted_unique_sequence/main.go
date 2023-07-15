package main

import (
	"fmt"
)

func main() {
	a := []int{1, 3, 4, 5, 6, 7, 8, 9, 10}
	prev := 0
	uniqueIndex := 0
	for _, val := range a {
		if prev != val {
			a[uniqueIndex] = val
			uniqueIndex++
		}
		prev = val
	}
	fmt.Println(a[:uniqueIndex])
}
