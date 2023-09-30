package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	done := make(chan struct{})
	gen := generator(&wg, done)

	wg.Add(2)

	go func(gen chan int) {
		defer wg.Done()
		for val := range gen {
			fmt.Println(val)
		}
	}(gen)

	time.Sleep(3 * time.Second)
	done <- struct{}{}
	close(done)

	wg.Wait()
}

func generator(wg *sync.WaitGroup, done chan struct{}) chan int {
	ch := make(chan int)
	go func(wg *sync.WaitGroup, ch chan int, done chan struct{}) {
		defer wg.Done()
		var i int
		for {
			select {
			case <-done:
				close(ch)
				return
			case <-time.After(500 * time.Millisecond):
				i++
				ch <- i
			}
		}
	}(wg, ch, done)
	return ch
}
