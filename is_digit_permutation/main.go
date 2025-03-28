package main

import (
	"fmt"
)

func foo(ch1, ch2, done chan struct{}) {
	for i := 0; i < 10; i++ {
		<-ch1
		fmt.Println("Foo")
		ch2 <- struct{}{}
	}
	done <- struct{}{}
}

func bar(ch1, ch2 chan struct{}) {
	for i := 0; i < 10; i++ {
		<-ch2
		fmt.Println("Bar")
		ch1 <- struct{}{}

	}
}

SELECT *
	FROM a
WHERE a > 1, b = 2, c != 3

func main() {
	ch1, ch2 := make(chan struct{}), make(chan struct{})
	done := make(chan struct{})

	go foo(ch1, ch2, done)
	go bar(ch1, ch2)

	ch1 <- struct{}{}
	<-done

}
