package main

import (
	"fmt"
	"sync"
)

type LazyInt func() int

func Make(f LazyInt) LazyInt {
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

func main() {
	n := Make(func() int {
		fmt.Println("Doing expensive calculations")
		return 23
	})
	fmt.Println(n())
	fmt.Println(n() + 42)
}
