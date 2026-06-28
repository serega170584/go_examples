package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type LazyFunc func() int

func generateLazy() LazyFunc {
	var o sync.Once
	var x int
	return func() int {
		o.Do(func() {
			x = rand.Intn(10000)
		})

		return x
	}
}

func main() {
	f := generateLazy()
	fmt.Println(f())
	fmt.Println(f())
}
