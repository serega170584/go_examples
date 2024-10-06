package main

import (
	"fmt"
	"sync"
)

func main() {
	in := make(chan int, 1000)
	for i := 0; i < 1000; i++ {
		in <- i
	}
	close(in)

	out := make(chan int)

	wg := &sync.WaitGroup{}
	wg.Add(4)
	for i := 0; i < 4; i++ {
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
