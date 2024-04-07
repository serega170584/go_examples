package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	in := make(chan int)
	out1 := make(chan int)
	out2 := make(chan int)
	out := []chan int{
		out1,
		out2,
	}

	go func(in chan int) {
		in <- 1
		in <- 2
		in <- 3
		close(in)
	}(in)

	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context, out1 chan int, out2 chan int) {
		for {
			select {
			case v := <-out1:
				fmt.Println(fmt.Sprintf("got value %d", v))
			case v := <-out2:
				fmt.Println(fmt.Sprintf("got value %d", v))
			case <-ctx.Done():
				break
			}
		}
	}(ctx, out1, out2)

	tee(ctx, in, out)

	time.Sleep(time.Millisecond)
	cancel()
}

func tee(ctx context.Context, in chan int, out []chan int) {
	for v := range in {
		for _, outCh := range out {
			go func(outCh chan int, v int) {
				select {
				case <-ctx.Done():
					break
				default:
					outCh <- v
					break
				}
			}(outCh, v)
		}
	}
}
