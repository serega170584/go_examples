package main

import (
	"fmt"
	"sync"
)

func main() {
	a := make(chan int, 4)
	a <- 1
	a <- 2
	a <- 3
	a <- 4
	close(a)
	b := make(chan int, 3)
	b <- 5
	b <- 6
	b <- 7
	close(b)
	c := make(chan int, 5)
	c <- 11
	c <- 12
	c <- 13
	c <- 14
	c <- 15
	close(c)

	for v := range mergeChannels(a, b, c) {
		fmt.Println(v)
	}
}

func mergeChannels(chList ...chan int) chan int {
	wg := &sync.WaitGroup{}
	out := make(chan int)
	for _, ch := range chList {
		wg.Add(1)
		for v := range ch {
			go func(wg *sync.WaitGroup, v int, out chan int) {
				defer wg.Done()
				out <- v
			}(wg, v, out)
		}
	}

	go func(wg *sync.WaitGroup, out chan int) {
		wg.Wait()
		close(out)
	}(wg, out)

	return out
}
