package main

import (
	"fmt"
	"strconv"
	"sync"
)

const workersCnt = 4

func main() {
	in := make(chan int, 10000)
	for i := 0; i < 10000; i++ {
		in <- i
	}
	close(in)

	output := make(chan string)

	wg := &sync.WaitGroup{}
	wg.Add(workersCnt)
	for i := 0; i < workersCnt; i++ {
		go func(wg *sync.WaitGroup, in chan int, output chan string, i int) {
			defer wg.Done()
			for v := range in {
				output <- strconv.Itoa(v) + ", task number: " + strconv.Itoa(i)
			}
		}(wg, in, output, i)
	}

	go func(wg *sync.WaitGroup, output chan string) {
		wg.Wait()
		close(output)
	}(wg, output)

	for v := range output {
		fmt.Println(v)
	}
}
