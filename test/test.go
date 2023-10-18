package main

import "fmt"

type I interface {
	test()
}

type S struct{}

func (s *S) test() {
	fmt.Println("asdasdasd")
}

func main() {
	var s S
	s.test()
}
