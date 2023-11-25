package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	wg := sync.WaitGroup{}

	input := make(chan int)
	output := make([]chan int, 2)
	output[0] = make(chan int)
	output[1] = make(chan int)

	done := make(chan struct{})

	wg.Add(4)
	go func(ctx context.Context, wg *sync.WaitGroup) {
		defer wg.Done()
		i := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Input closed")
				close(input)
				return
			case <-time.After(500 * time.Millisecond):
				fmt.Println("Put val to input ", i)
				input <- i
				i++
				fmt.Println("Put val to input ", i)
				input <- i
				i++
				fmt.Println("Put val to input ", i)
				input <- i
			}
		}
	}(ctx, &wg)

	go teeChannel(input, output, done, &wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		<-done
		for _, ch := range output {
			close(ch)
		}
	}(&wg)

	for ind, ch := range output {
		ch := ch
		ind := ind
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			for val := range ch {
				fmt.Println("Got from output index ", ind, " value ", val)
			}
		}(&wg)
	}

	wg.Wait()
}

func teeChannel(input chan int, output []chan int, done chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range input {
		for ind, outCh := range output {
			fmt.Println("Got from input val ", val, " to out ind ", ind)
			outCh <- val
		}
	}

	done <- struct{}{}
}
