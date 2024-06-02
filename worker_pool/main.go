package main

import (
	"fmt"
	"strconv"
	"sync"
)

const workersCnt = 4

func main() {
	in := make(chan int, 10000)
	out := make(chan string)
	wg := &sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		in <- i
	}
	close(in)

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(i int) {
			defer wg.Done()
			for v := range in {
				out <- "Result: task - " + strconv.Itoa(i) + ", value - " + strconv.Itoa(v)
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
