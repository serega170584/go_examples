package main

import "fmt"

type A struct {
	a int
}

func (a A) setA(v int) {
	fmt.Printf("%p\n", &a)
	a.a = v
}

// 19 15
// 9 * 5 + 9 * 10 + 10 * 5 + 10 * 10

// 28 = 2 * 16 + 8 * 1 = 32 + 8 = 40
// 40 = 4 * 16 + 0 = 64

// 48 = 4 * 16 + 8 = 72
// 50 = 5 * 16 + 0 = 80
func main() {
	var a A
	fmt.Printf("%p\n", &a)
	a.setA(1)
	a.setA(1)
	a.setA(1)
	a.setA(1)
	a.setA(1)
	a.setA(1)
}
