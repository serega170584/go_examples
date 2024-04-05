package main

import (
	"fmt"
	"sync"
)

func main() {
	semBar := make([]chan struct{}, 10)
	semFoo := make([]chan struct{}, 11)
	for i := 0; i < 10; i++ {
		semBar[i] = make(chan struct{})
		semFoo[i] = make(chan struct{})
	}
	semFoo[10] = make(chan struct{})
	go func() {
		semFoo[0] <- struct{}{}
		<-semFoo[10]
	}()
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-semFoo[i]
			fmt.Println("Foo")
			semBar[i] <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-semBar[i]
			fmt.Println("Bar")
			semFoo[i+1] <- struct{}{}
		}
	}()
	wg.Wait()
}
