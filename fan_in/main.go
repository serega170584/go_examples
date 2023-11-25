package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type payload struct {
	sender  string
	message int
}

func main() {
	wg := sync.WaitGroup{}
	done := make(chan struct{})

	producers := make([]<-chan payload, 2)
	wg.Add(2)
	producers[0] = makeProducer("Alice", &wg, done)
	producers[1] = makeProducer("Jack", &wg, done)

	fanIn := make(chan payload)
	wg.Add(2)
	var counter int64
	makeConsumer(producers, &wg, fanIn, &counter)

	go func() {
		for {
			if int(counter) == 2 {
				close(fanIn)
				return
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range fanIn {
			fmt.Println("Producer ", val.sender, " completed ", val.message)
		}
	}()

	time.Sleep(time.Second)
	close(done)
	wg.Wait()
}

func makeProducer(sender string, wg *sync.WaitGroup, done <-chan struct{}) <-chan payload {
	res := make(chan payload)

	go func() {
		defer wg.Done()
		i := 0
		for {
			select {
			case <-done:
				fmt.Println("Producer ", sender, " done")
				close(res)
				return
			case <-time.After(500 * time.Millisecond):
				fmt.Println("Produce ", i)
				res <- payload{sender: sender, message: i}
				i++
			}
		}
	}()

	return res
}

func makeConsumer(producers []<-chan payload, wg *sync.WaitGroup, fanIn chan<- payload, counter *int64) {
	for _, producerCh := range producers {
		producerCh := producerCh
		go func() {
			defer wg.Done()
			for val := range producerCh {
				fmt.Println("Consume producer ", val.sender, " value ", val.message)
				fanIn <- val
			}
			atomic.AddInt64(counter, 1)
		}()
	}
}
