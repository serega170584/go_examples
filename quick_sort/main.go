package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type f func() int

func generator(f1 f) f {
	a := 1
	var o sync.Once
	return func() int {
		o.Do(func() {
			a = f1()
		})

		return a
	}
}
func main() {
	f2 := generator(func() int {
		return rand.Intn(10000)
	})

	fmt.Println(f2())
	fmt.Println(f2())
}
