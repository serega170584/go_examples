package main

import "fmt"

func main() {
	var a interface{}
	test(a)
}

func test(a interface{}) {
	fmt.Println(a == nil)
}
