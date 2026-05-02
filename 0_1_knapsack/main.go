package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type f func(a int) int
func main() {
	a := make([]*A, 1)
	a[0].setB(1)
	fmt.Println(a[0].getB())
}

func generator() func(a int) int {
	var o sync.Once
	var b int
	return func(a int) int {
		o.Do(func() {
			 b = rand.Intn(10000 + a)
		})
		return b
	}
}

func


