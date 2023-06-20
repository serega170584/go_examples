package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	time.NewTicker(time.Second)
	cnt := 50
	batchSize := 10
	batches := make(chan struct{}, batchSize)

	go func() {
		for {
			for i := 0; i < batchSize; i++ {
				batches <- struct{}{}
			}
			time.Sleep(3 * time.Second)
		}
	}()

	for i := 0; i < cnt; i++ {
		<-batches
		fmt.Println(rand.Int())
	}
}
