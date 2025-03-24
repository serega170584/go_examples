package main

import "fmt"

func main() {
	a := make(chan int, 10)
	b := make(chan int, 5)
	for i := 0; i < 10; i++ {
		a <- i
	}
	close(a)
	b = a
	for v := range b {
		fmt.Println(v)
	}
}
