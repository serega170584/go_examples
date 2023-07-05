package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	cnt := 50
	batchSize := 10
	interval := make(chan struct{}, batchSize)
	res := make(chan int, cnt)
	go func() {
		ticker := time.NewTicker(3 * time.Second)
		for {
			select {
			case <-ticker.C:
				for i := 0; i < batchSize; i++ {
					interval <- struct{}{}
				}
			}
		}
	}()

	go func() {
		for i := 0; i < cnt; i++ {
			res <- RPCCallWithInterval(interval)
		}
		close(res)
	}()

	for val := range res {
		fmt.Println(val)
	}
}

func RPCCall() int {
	return rand.Intn(1000)
}

func RPCCallWithInterval(interval chan struct{}) int {
	<-interval
	return RPCCall()
}
