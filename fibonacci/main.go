package main

import (
	"fmt"
)

func main() {
	cnt := 50
	valCh := make(chan int)
	quit := make(chan struct{})
	res := make(chan struct{})
	go func() {
		for i := 0; i < cnt; i++ {
			fmt.Println(<-valCh)
		}
		quit <- struct{}{}
	}()
	fibonacci(valCh, quit, res)
	<-res
}

func fibonacci(valCh chan int, quit chan struct{}, res chan struct{}) {
	go func() {
		x, y := 0, 1
		for {
			select {
			case valCh <- x:
				x, y = y, x+y
			case <-quit:
				res <- struct{}{}
				return
			}
		}
	}()
}
