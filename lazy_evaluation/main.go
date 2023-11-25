package main

import (
	"fmt"
	"sync"
)

type LazyInt func() int

func main() {
	f := MakeLazyInt(func() int {
		return 23
	})
	fmt.Println(f())
	fmt.Println(f())
}

func MakeLazyInt(f LazyInt) LazyInt {
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
