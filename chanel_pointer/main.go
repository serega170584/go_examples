package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan *int, 1)
	var a *int
	var b int
	a = &b
	ch <- a
	go func() {
		*a = 1
	}()
	time.Sleep(3 * time.Second)
	c := <-ch
	fmt.Println(*c)
}
