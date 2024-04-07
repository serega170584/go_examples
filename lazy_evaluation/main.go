package main

import (
	"fmt"
	"sync"
)

type LazyFunc func() int

func main() {
	f := makeFunc(func() int {
		return 23
	})
	fmt.Println(f())
}

func makeFunc(f LazyFunc) LazyFunc {
	var v int
	var once sync.Once
	return func() int {
		once.Do(func() {
			v = f()
			f = nil
		})
		return v
	}
}
