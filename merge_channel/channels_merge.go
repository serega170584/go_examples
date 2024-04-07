package main

import "fmt"

func main() {
	a := make(chan int, 3)
	a <- 1
	a <- 2
	a <- 3
	close(a)
	b := make(chan int, 4)
	b <- 1
	b <- 2
	b <- 3
	b <- 4
	close(b)
	c := make(chan int, 5)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	close(c)

	out := mergeChannels(a, b, c)

	for v := range out {
		fmt.Println(v)
	}
}

func mergeChannels(chList ...chan int) chan int {
	out := make(chan int)
	sem := make([]chan struct{}, len(chList))
	for i := range sem {
		sem[i] = make(chan struct{})
	}

	for i, ch := range chList {
		go func(ch chan int, out chan int, sem []chan struct{}, i int) {
			for v := range ch {
				out <- v
			}
			sem[i] <- struct{}{}
		}(ch, out, sem, i)
	}

	go func(sem []chan struct{}, out chan int) {
		for _, v := range sem {
			<-v
		}
		close(out)
	}(sem, out)

	return out
}
