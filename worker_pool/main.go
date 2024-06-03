package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	in := make(chan int, 1000)
	for i := 0; i < 1000; i++ {
		in <- i
	}
	close(in)

	out := make(chan string)

	wg := &sync.WaitGroup{}
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func(i int) {
			defer wg.Done()
			for v := range in {
				out <- "Result: " + strconv.Itoa(v) + "task: " + strconv.Itoa(i)
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	for v := range out {
		fmt.Println(v)
	}
}
