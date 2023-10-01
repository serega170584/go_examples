package main

import "fmt"

// write to res in cycle, read from res in cycle of main
// to do read from res in cycle finite we have to close res
// when close res?
//   - after working out all input channels => after working out channel 1, channel 2, ...
//
// Obviuosly we need count complteted works wuth input channels
// for that we use channel count of empty struct
// after doing input channel, write struct to counter
// in loop from i = 0 to count of input channels read from counter channel
// after having done cycle we can close res channel so we unlock main goroutine
// also we close counter after cycle
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
