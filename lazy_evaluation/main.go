package main

import (
	"fmt"
	"sync"
)

type LazyInt func() int

func main() {
	a := MakeLazyInt(func() int { return 1 })
	fmt.Println(a())
	fmt.Println(a())
}

func MakeLazyInt(f LazyInt) LazyInt {
	var o sync.Once
	res := 0
	return func() int {
		o.Do(func() {
			res = f()
			f = nil
		})
		return res
	}
}
