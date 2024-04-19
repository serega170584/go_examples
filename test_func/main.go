package main

import (
	"fmt"
	"sync"
)

type testFunc func(a int) int

func main() {
	resFunc := makeTestFunc(func(a int) int { return a + 3 })
	fmt.Println(resFunc(3))
	fmt.Println(resFunc(3))
}

func makeTestFunc(f testFunc) testFunc {
	var once sync.Once
	var res int
	return func(a int) int {
		once.Do(func() {
			res = f(a)
			f = nil
		})

		return res
	}
}
