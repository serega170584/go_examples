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
	res := mergeChannels(ch1, ch2, ch3)
	for val := range res {
		fmt.Println(val)
	}
}

func mergeChannels(chList ...chan int) chan int {
	res := make(chan int)
	go func() {
		for _, chValues := range chList {
			for val := range chValues {
				res <- val
			}
		}
		close(res)
	}()
	return res
}
