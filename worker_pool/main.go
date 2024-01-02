package main

import (
	"fmt"
	"sync"
	"time"
)

const jobsCnt = 10
const workersCnt = 3

func main() {
	jobs := make(chan int, jobsCnt)
	go func(jobs chan<- int) {
		for i := 0; i < jobsCnt; i++ {
			fmt.Println("Initialized job ", i)
			jobs <- i
		}
		close(jobs)
	}(jobs)

	wg := &sync.WaitGroup{}
	wg.Add(workersCnt)
	for i := 0; i < workersCnt; i++ {
		go worker(jobs, wg)
	}
	wg.Wait()
}

func worker(jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range jobs {
		fmt.Println("Processed value ", val)
		time.Sleep(500 * time.Millisecond)
	}
}
