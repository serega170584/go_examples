package main

import "fmt"

func main() {
	ch1 := make(chan int, 4)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	ch1 <- 4
	close(ch1)
	ch2 := make(chan int, 4)
	ch2 <- 5
	ch2 <- 6
	ch2 <- 7
	ch2 <- 8
	close(ch2)
	ch3 := make(chan int, 3)
	ch3 <- 9
	ch3 <- 10
	ch3 <- 11
	close(ch3)
	ch4 := make(chan int, 2)
	ch4 <- 12
	ch4 <- 13
	close(ch4)

	output := mergeChannels(ch1, ch2, ch3, ch4)

	for v := range output {
		fmt.Println(v)
	}
}

func mergeChannels(chList ...chan int) chan int {
	output := make(chan int)
	sem := make([]chan struct{}, len(chList))

	for i := range sem {
		sem[i] = make(chan struct{})
	}

	for i, ch := range chList {
		go func(ch chan int, i int, output chan int, sem []chan struct{}) {
			for v := range ch {
				output <- v
			}
			sem[i] <- struct{}{}
		}(ch, i, output, sem)
	}

	go func(sem []chan struct{}, output chan int) {
		for _, ch := range sem {
			<-ch
		}
		close(output)
	}(sem, output)

	return output
}
