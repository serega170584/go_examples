package main

import (
	"fmt"
	"math/rand"
)

type sema chan struct{}

func main() {
	s := make(sema)
	for i := 0; i < 10; i++ {
		go func(s sema) {
			fmt.Println(rand.Intn(1000))
			s.inc(1)
		}(s)
	}

	s.dec(10)
}

func (s sema) inc(k int) {
	for i := 0; i < k; i++ {
		s <- struct{}{}
	}
}

func (s sema) dec(k int) {
	for i := 0; i < k; i++ {
		<-s
	}
}
