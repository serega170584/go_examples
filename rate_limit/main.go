package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	cnt := 50
	batchSize := 10
	resCh := make(chan int, batchSize)
	interval := make(chan struct{}, batchSize)

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
			resCh <- RPCCallWithInterval(interval)
		}
		close(resCh)
	}()

	for val := range resCh {
		fmt.Println(val)
	}
}

func RPCCallWithInterval(interval chan struct{}) int {
	<-interval
	return RPCCall()
}

func RPCCall() int {
	return rand.Intn(1000)
}
