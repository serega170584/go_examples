package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		a := make([]int, 0)
		defer wg.Done()
		select {
		case <-ch1:
			fmt.Println("1")
		case <-ch2:
			a[0] = 1
		}
	}()

	ch1 <- 123
	ch2 <- 123
	close(ch1)
	close(ch2)
	wg.Wait()
}
