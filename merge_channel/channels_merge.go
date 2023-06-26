package main

import "fmt"

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

	output := mergeChannels(ch1, ch2, ch3)

	for val := range output {
		fmt.Println(val)
	}
}

func mergeChannels(in ...chan int) chan int {
	var cnt int
	for _, vals := range in {
		cnt += len(vals)
	}

	counter := make(chan struct{})
	output := make(chan int)
	for _, vals := range in {
		vals := vals
		go func() {
			for val := range vals {
				output <- val
				counter <- struct{}{}
			}
		}()
	}

	go func() {
	loop:
		for {
			<-counter
			cnt--
			if cnt == 0 {
				break loop
			}
		}
		close(output)
	}()

	return output
}
