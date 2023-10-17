package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func printFoo(handledFooCnt *int64, handledBarCnt *int64) {
	for i := 0; i < 10; i++ {
		for *handledFooCnt != *handledBarCnt {
			continue
		}
		fmt.Println("Foo")
		atomic.AddInt64(handledFooCnt, 1)
	}
}

func printBar(handledFooCnt *int64, handledBarCnt *int64, ch <-chan struct{}) {
	for i := 0; i < 10; i++ {
		for *handledFooCnt != *handledBarCnt+1 {
			continue
		}
		fmt.Println("Bar")
		atomic.AddInt64(handledBarCnt, 1)
	}
	<-ch
}

func main() {
	t := time.Now()
	var handledFooCnt, handledBarCnt *int64
	var a, b int64
	handledFooCnt = &a
	handledBarCnt = &b
	ch := make(chan struct{})
	go printFoo(handledFooCnt, handledBarCnt)
	go printBar(handledFooCnt, handledBarCnt, ch)
	ch <- struct{}{}
	close(ch)
	fmt.Println(time.Since(t))
}
