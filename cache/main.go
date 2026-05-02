package main

import (
	"fmt"
	"sync"
)

func merge(chList ...chan int) chan int {
	var wg sync.WaitGroup
	result := make(chan int)

	for _, ch := range chList {
		wg.Go(func() {
			for value := range ch {
				result <- value
			}
		})
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	return result
}

func main() {
	ch1 := make(chan int, 2)
	ch1 <- 1
	ch1 <- 2
	close(ch1)
	ch2 := make(chan int, 2)
	ch2 <- 1
	ch2 <- 2
	close(ch2)

	out := merge(ch1, ch2)

	for v := range out {
		fmt.Println(v)
	}
}
