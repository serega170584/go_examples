package main

import "fmt"

func main() {
	in := make(chan int)
	output := make(chan int, 4)

	go func(in, output chan int) {
		for v := range in {
			select {
			case output <- v:
				continue
			default:
				<-output
				output <- v
			}
		}
		close(output)
	}(in, output)

	for i := 0; i < 10000; i++ {
		in <- i
	}
	close(in)

	for v := range output {
		fmt.Println(v)
	}

}
