package main

import "fmt"

func main() {
	out := make(chan int, 4)
	in := make(chan int)
	go func() {
		for v := range in {
			select {
			case out <- v:
				continue
			case <-out:
				out <- v
			}
		}
		close(out)
	}()

	for i := 0; i < 10000; i++ {
		in <- i
	}
	close(in)

	for v := range out {
		fmt.Println(v)
	}
}
