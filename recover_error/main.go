package main

import "fmt"

func main() {
	defer handler()
	a := make([]int, 0)
	a[1] = 123

	go func() {
		defer handler()
		a := make([]int, 0)
		a[1] = 123
	}()
}

func handler() {
	e := recover()
	err, _ := e.(error)
	fmt.Println(err)
}
