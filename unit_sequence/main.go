package main

import (
	"fmt"
	"math/rand"
)

func main() {
	cnt := 20
	a := make([]int, cnt)
	counter := 0
	maxLength := 0
	for i := 0; i < cnt; i++ {
		a[i] = rand.Intn(2)
	}

	fmt.Println(a)

	a = append(a, 3)
	for _, val := range a {
		if val == 1 {
			counter++
		} else if counter > maxLength {
			maxLength = counter
		}
		if val != 1 {
			counter = 0
		}
	}
	fmt.Println(maxLength)
}
