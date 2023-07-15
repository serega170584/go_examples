package main

import (
	"fmt"
	"math/rand"
)

func main() {
	cnt := 10
	a := make([]int, cnt)
	b := make([]int, cnt)
	uniqueA := make(map[int]struct{}, cnt)
	intersection := make(map[int]int, cnt)
	for i := 0; i < cnt; i++ {
		a[i] = rand.Intn(20)
		b[i] = rand.Intn(20) + 10
	}
	fmt.Println(a)
	fmt.Println(b)

	for _, val := range a {
		uniqueA[val] = struct{}{}
	}

	for _, val := range b {
		if _, ok := uniqueA[val]; ok {
			intersection[val] = val
		}
	}

	fmt.Println(intersection)
}
