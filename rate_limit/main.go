package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	batchSize := 10
	cnt := 50
	interval := make(chan struct{})
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

	for i := 0; i < cnt; i++ {
		fmt.Println(RPCCallWithInterval(interval))
	}
}

func RPCCall() int {
	return rand.Intn(1000)
}

func RPCCallWithInterval(interval chan struct{}) int {
	<-interval
	return RPCCall()
}
