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
	ch2 := make(chan int, 3)
	ch2 <- 5
	ch2 <- 6
	ch2 <- 7
	close(ch2)
	ch3 := make(chan int, 2)
	ch3 <- 8
	ch3 <- 9
	close(ch3)

	out := merge(ch1, ch2, ch3)

	for v := range out {
		fmt.Println(v)
	}
}

func merge(chList ...chan int) chan int {
	res := make(chan int)
	var wg sync.WaitGroup
	for _, ch := range chList {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range ch {
				res <- v
			}
		}()
	}
	go func() {
		wg.Wait()
		close(res)
	}()
	return res
}
