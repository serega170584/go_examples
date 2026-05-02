package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	f := getFunc()
	fmt.Println(f())
	fmt.Println(f())
}

func getFunc() func() int {
	var o sync.Once
	var x int
	return func() int {
		o.Do(func() {
			x = rand.Intn(2000)
		})

		return x
	}
}
