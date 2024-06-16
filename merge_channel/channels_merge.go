package main

import (
	"fmt"
	"sync"
)

func main() {
	a := make(chan int, 3)
	a <- 1
	a <- 2
	a <- 3
	close(a)
	b := make(chan int, 2)
	b <- 4
	b <- 5
	close(b)
	c := make(chan int, 4)
	c <- 6
	c <- 7
	c <- 8
	c <- 9
	close(c)

	out := merge(a, b, c)

	for v := range out {
		fmt.Println(v)
	}
}

func merge(chList ...chan int) chan int {
	out := make(chan int)
	l := len(chList)
	wg := &sync.WaitGroup{}
	wg.Add(l)
	for _, ch := range chList {
		go func(ch chan int) {
			defer wg.Done()
			for v := range ch {
				out <- v
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
