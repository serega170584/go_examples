package main

import "fmt"

const (
	_ = iota
	TEST1
	TEST2
	TEST3
	TEST4 = iota
	TEST5
)

func main() {
	fmt.Println(TEST1)
}
