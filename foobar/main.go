package main

import "fmt"

func main() {
	semBar := make([]chan struct{}, 10)
	semFoo := make([]chan struct{}, 11)
	for i := 0; i < 10; i++ {
		semBar[i] = make(chan struct{})
		semFoo[i] = make(chan struct{})
	}
	semFoo[10] = make(chan struct{})

	go func(semFoo []chan struct{}) {
		semFoo[0] <- struct{}{}
	}(semFoo)

	go func(semFoo []chan struct{}, semBar []chan struct{}) {
		for i := 0; i < 10; i++ {
			<-semFoo[i]
			fmt.Println("Foo")
			semBar[i] <- struct{}{}
		}
	}(semFoo, semBar)

	go func(semFoo []chan struct{}, semBar []chan struct{}) {
		for i := 0; i < 10; i++ {
			<-semBar[i]
			fmt.Println("Bar")
			semFoo[i+1] <- struct{}{}
		}
	}(semFoo, semBar)

	<-semFoo[10]
}
