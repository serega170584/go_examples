package main

import "fmt"

func printFoo(ch chan struct{}, exch chan struct{}, res chan struct{}) {
	for i := 0; i < 10; i++ {
		<-ch
		fmt.Print("Foo")
		exch <- struct{}{}
	}
	<-ch
	<-res
}

func printBar(ch chan struct{}, exch chan struct{}) {
	for i := 0; i < 10; i++ {
		<-exch
		fmt.Print("Bar\n")
		ch <- struct{}{}
	}
}

func main() {
	ch := make(chan struct{})
	exch := make(chan struct{})
	res := make(chan struct{})
	go printFoo(ch, exch, res)
	go printBar(ch, exch)
	ch <- struct{}{}
	res <- struct{}{}
}
