package main

import "fmt"

func main() {
	a := make(chan int, 3)
	a <- 1
	a <- 2
	a <- 3
	close(a)
	b := make(chan int, 2)
	b <- 2
	b <- 3
	close(b)
	c := make(chan int, 4)
	c <- 4
	c <- 5
	c <- 6
	c <- 7
	close(c)

	o := mergeChannels(a, b, c)

	for v := range o {
		fmt.Println(v)
	}
}

func mergeChannels(chList ...chan int) chan int {
	output := make(chan int)
	sem := make([]chan struct{}, len(chList))

	for i := range chList {
		sem[i] = make(chan struct{})
	}

	for i, ch := range chList {
		go func(ch chan int, output chan int, i int, sem []chan struct{}) {
			for v := range ch {
				output <- v
			}
			sem[i] <- struct{}{}
		}(ch, output, i, sem)
	}

	go func(sem []chan struct{}, output chan int) {
		for _, v := range sem {
			<-v
		}
		close(output)
	}(sem, output)

	return output
}
