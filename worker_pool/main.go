package main

import (
	"fmt"
	"sync"
)

func main() {
	in := make(chan int, 1000)
	out := make(chan int)

	wg := &sync.WaitGroup{}
	wg.Add(5)

	for i := 0; i < 1000; i++ {
		in <- i
	}
	close(in)

	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			for v := range in {
				out <- v
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	for v := range out {
		fmt.Println(v)
	}
}
