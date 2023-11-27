package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

const TOTAL = 10
const SemaphoreCnt = 3

type Semaphore chan struct{}

func NewSemaphore(cnt int) Semaphore {
	return make(Semaphore, cnt)
}

func (s Semaphore) Acquire() {
	s <- struct{}{}
}

func (s Semaphore) Release() {
	<-s
}

func process(v int) {
	fmt.Println(time.Now().Format("Mon Jan _2 15:04:04 2006"), " running task", v)
}

func main() {
	s := NewSemaphore(SemaphoreCnt)

	done := make(chan struct{})

	var counter int64

	for i := 0; i < TOTAL; i++ {
		s.Acquire()
		go func(val int) {
			defer s.Release()
			process(val)
			atomic.AddInt64(&counter, 1)
			if int(counter) == TOTAL {
				close(done)
			}
		}(i)
	}

	<-done
}
