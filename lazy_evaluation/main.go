package main

import (
	"fmt"
	"sync"
)

type LazyInt func() int

func main() {
	b := generateLazyInt(func() int {
		return 123
	})
	fmt.Println(b())
	fmt.Println(b())
}

func generateLazyInt(f LazyInt) LazyInt {
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
