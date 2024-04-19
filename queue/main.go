package main

import (
	"fmt"
	"sync"
)

const messagesCnt = 10
const queueSize = 3

func main() {
	wg := &sync.WaitGroup{}
	queue := make(chan struct{}, queueSize)
	for i := 0; i < messagesCnt; i++ {
		fmt.Println("Try ", i)
		handle(queue, i, wg)
	}
	wg.Wait()
}

func handle(queue chan struct{}, i int, wg *sync.WaitGroup) {
	queue <- struct{}{}
	wg.Add(1)
	go func(wg *sync.WaitGroup, i int) {
		defer wg.Done()
		fmt.Println("Start ", i)
		fmt.Println("Finish ", i)
		<-queue
	}(wg, i)
}
