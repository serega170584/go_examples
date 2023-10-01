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

	resCh := mergeChannels(ch1, ch2, ch3)

	for val := range resCh {
		fmt.Println(val)
	}
}

func mergeChannels(chList ...chan int) chan int {
	res := make(chan int)
	cnt := len(chList)
	counterCh := make(chan struct{})
	for _, ch := range chList {
		ch := ch
		go func() {
			for val := range ch {
				res <- val
			}
			counterCh <- struct{}{}
		}()
	}

	go func() {
		for i := 0; i < cnt; i++ {
			<-counterCh
		}
		close(counterCh)
		close(res)
	}()

	return res
}
