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
	b := make(chan int, 4)
	b <- 4
	b <- 5
	b <- 6
	b <- 7
	close(b)
	c := make(chan int, 2)
	c <- 8
	c <- 9
	close(c)

	output := merge(a, b, c)
	for v := range output {
		fmt.Println(v)
	}
}

func merge(chList ...chan int) chan int {
	output := make(chan int)
	wg := &sync.WaitGroup{}
	for _, in := range chList {
		wg.Add(1)
		go func(in chan int, output chan int, wg *sync.WaitGroup) {
			defer wg.Done()
			for v := range in {
				output <- v
			}
		}(in, output, wg)
	}

	go func(wg *sync.WaitGroup) {
		wg.Wait()
		close(output)
	}(wg)

	return output
}
