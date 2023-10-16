package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	input := make(chan string)
	output1 := make(chan string)
	output2 := make(chan string)
	outputs := []chan<- string{output1, output2}

	go func() {
		input <- "A"
		input <- "B"
		input <- "C"
		close(input)
	}()

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case elem := <-output1:
				fmt.Println("Got value: ", elem)
			case elem := <-output2:
				fmt.Println("Got value ", elem)
			case <-ctx.Done():
				return
			}
		}
	}()

	teeChannel(ctx, input, outputs)

	time.Sleep(time.Second)
	cancel()
}

func teeChannel(ctx context.Context, input <-chan string, outputs []chan<- string) {
	for elem := range input {
		elem := elem
		for _, out := range outputs {
			out := out
			go func() {
				select {
				case <-ctx.Done():
					return
				case out <- elem:
					return
				}
			}()
		}
	}
}
