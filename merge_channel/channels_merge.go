package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	t := time.Now()
	defer func() {
		fmt.Println(time.Since(t).Microseconds())
	}()
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
	out := make(chan int)
	wg := &sync.WaitGroup{}
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
