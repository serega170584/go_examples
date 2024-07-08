package main

import "fmt"

func main() {
	fooSem := make(chan struct{})
	barSem := make(chan struct{})
	resSem := make(chan struct{})

	go foo(fooSem, barSem)
	go bar(fooSem, barSem, resSem)

	<-resSem
}

func foo(fooSem chan struct{}, barSem chan struct{}) {
	for i := 0; i < 10; i++ {
		if i != 0 {
			<-fooSem
		}
		fmt.Println("Foo")
		barSem <- struct{}{}
	}
}

func bar(fooSem chan struct{}, barSem chan struct{}, resSem chan struct{}) {
	for i := 0; i < 10; i++ {
		<-barSem
		fmt.Println("Bar")
		if i == 9 {
			resSem <- struct{}{}
		} else {
			fooSem <- struct{}{}
		}
	}
}
