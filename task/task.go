package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	fmt.Printf("%d\n", len(ch))
	fmt.Printf("%d\n", cap(ch))

	ch <- 1
	ch <- 2
	fmt.Printf("%d\n", len(ch))
	fmt.Printf("%d\n", cap(ch))

}
