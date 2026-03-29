package main

import (
	"fmt"
	"sync"
)

func merge(chList ...chan int) chan int {
	result := make(chan int)
	go func() {
		defer close(result)
		wg := sync.WaitGroup{}
		for _, ch := range chList {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for cur := range ch {
					result <- cur
				}
			}()
		}
		wg.Wait()
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
