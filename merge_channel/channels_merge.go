package main

import "fmt"

func main() {
	ch1 := make(chan int, 3)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	close(ch1)
	ch2 := make(chan int, 3)
	ch2 <- 4
	ch2 <- 5
	ch2 <- 6
	close(ch2)
	ch3 := make(chan int, 3)
	ch3 <- 7
	ch3 <- 8
	ch3 <- 9
	close(ch3)

	output := mergeChannels(ch1, ch2, ch3)
	for val := range output {
		fmt.Println(val)
	}
}

func mergeChannels(chList ...chan int) chan int {
	var cnt int
	for _, ch := range chList {
		cnt += len(ch)
	}

	output := make(chan int, cnt)
	go func() {
		for _, vals := range chList {
			for val := range vals {
				output <- val
			}
		}
		close(output)
	}()

	return output
}
