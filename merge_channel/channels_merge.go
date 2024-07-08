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
	b := make(chan int, 2)
	b <- 5
	b <- 6
	close(b)
	c := make(chan int, 3)
	c <- 7
	c <- 8
	c <- 9
	close(c)

	out := mergeChannels(a, b, c)
	for v := range out {
		fmt.Println(v)
	}
}

func mergeChannels(chList ...chan int) chan int {
	l := len(chList)
	out := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(l)
	for _, ch := range chList {
		for v := range ch {
			go func(v int) {
				defer wg.Done()
				out <- v
			}(v)
		}
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
