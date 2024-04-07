package main

import "fmt"

func main() {
	in := make(chan int)
	out := make(chan int, 4)

	go func() {
		for v := range in {
			select {
			case out <- v:
			default:
				<-out
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
