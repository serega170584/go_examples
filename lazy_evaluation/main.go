package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type LazyFunc func() int

func main() {
	f := generate(func() int { return rand.Intn(10000) })
	fmt.Println(f())
	fmt.Println(f())
}

func generate(f LazyFunc) LazyFunc {
	var o sync.Once
	var a int
	return func() int {
		o.Do(func() {
			a = f()
			f = nil
		})

		return a
	}
}
