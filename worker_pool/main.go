package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

const WorkerCnt = 3
const MessagesCnt = 10

func main() {
	start := time.Now()

	jobs := make(chan int, MessagesCnt)
	go func() {
		for i := 0; i < MessagesCnt; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	var counter int64

	done := make(chan struct{})

	for i := 0; i < WorkerCnt; i++ {
		go worker(jobs, &counter)
	}

	go func() {
		for {
			if int(counter) == MessagesCnt {
				close(done)
				return
			}
		}
	}()

	<-done
	fmt.Println("Duration: ", time.Since(start))
}

func worker(jobs chan int, counter *int64) {
	for val := range jobs {
		fmt.Println(time.Now().Format("Mon Jan _2 15:04:04 2006"), " got value ", val)
		atomic.AddInt64(counter, 1)
	}
}
