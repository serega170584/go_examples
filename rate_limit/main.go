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
	res := make(chan int, batchSize)
	go func() {
		ticker := time.NewTicker(5 * time.Second)
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
			res <- RPCWithInterval(interval)
		}
		close(res)
		close(interval)
	}()

	for val := range res {
		fmt.Println(val)
	}
}

func RPCWithInterval(interval chan struct{}) int {
	<-interval
	return RPCCall()
}

func RPCCall() int {
	return rand.Intn(1000)
}
