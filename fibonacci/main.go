package main

import "fmt"

func main() {
	cnt := 50
	ch := make(chan int)
	quit := make(chan struct{})
	go func() {
		for i := 0; i < cnt; i++ {
			fmt.Println(<-ch)
		}
		quit <- struct{}{}
	}()
	fibonacci(ch, quit)
}

func fibonacci(ch chan int, quit chan struct{}) {
	x, y := 0, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-quit:
			return
		}
	}
}
