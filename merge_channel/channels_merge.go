package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int, 4)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	ch1 <- 4
	close(ch1)
	ch2 := make(chan int, 2)
	ch2 <- 5
	ch2 <- 6
	close(ch2)
	ch3 := make(chan int, 3)
	ch3 <- 7
	ch3 <- 8
	ch3 <- 9
	close(ch3)

	out := merge(ch1, ch2, ch3)

	for v := range out {
		fmt.Println(v)
	}
}

func merge(chList ...chan int) chan int {
	out := make(chan int)
	wg := sync.WaitGroup{}
	l := len(chList)
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
