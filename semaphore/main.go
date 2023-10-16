package main

import (
	"fmt"
	"time"
)

type semaphore chan struct{}

func newSemaphore(n int) semaphore {
	return make(semaphore, n)
}

func (s semaphore) Acquire(n int) {
	for i := 0; i < n; i++ {
		s <- struct{}{}
	}
}

func (s semaphore) Release(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}

const N = 3
const TOTAL = 10

func main() {
	s := newSemaphore(N)
	done := make(chan struct{})

	for i := 0; i < TOTAL; i++ {
		s.Acquire(1)
		go func(i int) {
			defer s.Release(1)
			process(i)
			if i == TOTAL-1 {
				done <- struct{}{}
			}
		}(i)
	}

	<-done
	close(done)
}

func process(i int) {
	fmt.Println("Got value ", i, " date: ", time.Now())
}
