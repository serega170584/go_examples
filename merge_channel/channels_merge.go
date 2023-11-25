package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	ch1 := make(chan int, 3)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	close(ch1)

	ch2 := make(chan int, 2)
	ch2 <- 4
	ch2 <- 5
	close(ch2)

	ch3 := make(chan int, 4)
	ch3 <- 6
	ch3 <- 7
	ch3 <- 8
	ch3 <- 9
	close(ch3)

	res := mergeChannels(ch1, ch2, ch3)
	for val := range res {
		fmt.Println(val)
	}
}

func mergeChannels(chList ...chan int) <-chan int {
	res := make(chan int)
	var counter int64

	for _, ch := range chList {
		ch := ch
		go func() {
			for val := range ch {
				res <- val
			}
			atomic.AddInt64(&counter, 1)
		}()
	}

	go func() {
		for {
			if int(counter) == len(chList) {
				close(res)
				return
			}
		}
	}()

	return res
}
