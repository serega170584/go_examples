package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int, 3)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	ch2 := make(chan int, 3)
	ch2 <- 4
	ch2 <- 5
	ch2 <- 6
	ch3 := make(chan int, 3)
	ch3 <- 7
	ch3 <- 8
	ch3 <- 9

	ch := mergeChannels(ch1, ch2, ch3)
	for val := range ch {
		fmt.Println(val)
	}
}

func mergeChannels(chList ...chan int) chan int {
	var cnt int
	for _, vals := range chList {
		cnt += len(vals)
	}
	tmp := make(chan int, cnt)
	output := make(chan int, cnt)

	for _, ch := range chList {
		ch := ch
		go func() {
			for val := range ch {
				tmp <- val
			}
		}()
	}

	go func() {
		for i := 0; i < cnt; i++ {
			output <- <-tmp
		}
		close(output)
		close(tmp)
	}()

	return output
}
