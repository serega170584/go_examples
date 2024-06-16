package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	in := make(chan int, 1000)
	out := make(chan string)

	for i := 0; i < 1000; i++ {
		in <- i
	}
	close(in)

	wg := &sync.WaitGroup{}
	wg.Add(4)

	for i := 0; i < 4; i++ {
		go func(i int) {
			defer wg.Done()
			for v := range in {
				out <- "Task: " + strconv.Itoa(i) + ", value: " + strconv.Itoa(v)
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
