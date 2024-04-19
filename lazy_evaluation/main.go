package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type LazyInt func() int

func a() int {
	return rand.Intn(100)
}

func main() {
	b := makeLazyInt(a)
	c := b()
	fmt.Println(c)
	d := b()
	fmt.Println(d)
}

func makeLazyInt(f LazyInt) LazyInt {
	var v int
	var o sync.Once
	return func() int {
		o.Do(func() {
			v = f()
			f = nil
		})
		return v
	}
}
