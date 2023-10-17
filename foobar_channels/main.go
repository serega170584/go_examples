package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	in := make(chan struct{})
	out := make(chan struct{})
	done := make(chan struct{})
	go printFoo(in, out)
	go printBar(out, in, done)
	in <- struct{}{}
	<-done
	close(in)
	close(out)
	fmt.Println(time.Since(t))
}

func printFoo(in <-chan struct{}, out chan<- struct{}) {
	for i := 0; i < 10; i++ {
		<-in
		fmt.Println("Foo")
		out <- struct{}{}
	}
}

func printBar(in <-chan struct{}, out chan<- struct{}, done chan<- struct{}) {
	for i := 0; i < 10; i++ {
		<-in
		fmt.Println("Bar")
		if i == 9 {
			break
		}
		out <- struct{}{}
	}
	done <- struct{}{}
}
