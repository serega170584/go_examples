package main

import "fmt"

type TestInterface interface {
	A()
}

type A struct {
}

func (*A) A() {}

func main() {
	var a interface{}
	var b *A
	a = b
	//var a TestInterface
	//a = &A{}
	//a = 1
	//b := a.(*A)
	fmt.Println(a == nil)
	fmt.Println(a.(*A) == nil)
}
