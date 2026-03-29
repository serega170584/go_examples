package main

import (
	"fmt"
	"sync"
)

type LazyFunc func() int

func New(f LazyFunc) LazyFunc {
	x := 0
	var o sync.Once
	return func() int {
		o.Do(func() {
			x = f()
			f = nil
		})
		return x
	}
}

func main() {
	f := New(func() int {
		return 123
	})
	fmt.Println(f())
}
