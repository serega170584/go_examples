package main

import (
	"fmt"
	"sync"
)

type LazyInt func() int

func makeLazyInt(f LazyInt) LazyInt {
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

func main() {
	b := makeLazyInt(func() int {
		return 123
	})
	fmt.Println(b())
	fmt.Println(b())
}
