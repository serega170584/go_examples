package main

import (
	"fmt"
	"github.com/google/uuid"
)

//go:noinline
func main() {
	fmt.Println(uuid.NewUUID())
	//a := make([]uint64, 2<<20)
	var b, c, d int
	b = 12
	c = 34
	d = 56
	//println(a)
	println(&b)
	println(&c)
	e()
	e()
	//e()
	//e()
	var f int
	println(&d)
	println(&f)
}

func e() {
	a := 1
	println(&a)
}
