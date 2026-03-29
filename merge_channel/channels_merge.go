package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int, 2)
	ch1 <- 1
	ch1 <- 2
	close(ch1)
	ch2 := make(chan int, 3)
	ch2 <- 3
	ch2 <- 4
	ch2 <- 5
	close(ch2)
	ch3 := make(chan int, 4)
	ch3 <- 6
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
	var wg sync.WaitGroup
	wg.Add(len(chList))
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
