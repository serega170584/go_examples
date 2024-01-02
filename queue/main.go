package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const queueCnt = 3
const messagesCnt = 10

func main() {
	queue := make(chan struct{}, queueCnt)
	wg := &sync.WaitGroup{}
	for i := 0; i < messagesCnt; i++ {
		num := rand.Intn(1000)
		fmt.Println("Created num ", num)
		process(num, queue, wg)
	}
	wg.Wait()
	close(queue)
}

func process(val int, queue chan struct{}, wg *sync.WaitGroup) {
	queue <- struct{}{}

	wg.Add(1)
	go func(wg *sync.WaitGroup, val int) {
		defer wg.Done()
		fmt.Println("Processing start ", val)
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Processing finish", val)
		<-queue
	}(wg, val)
}
