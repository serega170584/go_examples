package main

import (
	"fmt"
	"sync"
)

type LazyInt func() int

func main() {
	f := Make(func() int {
		return 23
	})
	fmt.Println(f())
	fmt.Println(f() + 23)
}

func Make(f LazyInt) LazyInt {
	var once sync.Once
	var v int
	return func() int {
		once.Do(
			func() {
				v = f()
				f = nil
			})
		return v
	}
}
