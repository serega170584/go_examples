package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type payload struct {
	sender string
	value  int
}

func main() {
	done := make(chan struct{})
	producers := make([]chan payload, 2)

	wg := sync.WaitGroup{}

	wg.Add(2)
	producers[0] = makeProducer("Alice", done, &wg)
	producers[1] = makeProducer("Max", done, &wg)

	wg.Add(2)
	fanIn := make(chan payload)
	var counter int64
	consumer(producers, fanIn, &wg, &counter)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for val := range fanIn {
			fmt.Println("Got sender: ", val.sender, " value: ", val.value)
		}
	}(&wg)

	go func(counter *int64) {
		for {
			if int(*counter) == 2 {
				close(fanIn)
				return
			}
		}
	}(&counter)

	time.Sleep(time.Second)
	close(done)

	wg.Wait()
}

func makeProducer(sender string, done <-chan struct{}, wg *sync.WaitGroup) chan payload {
	output := make(chan payload)
	go func(wg *sync.WaitGroup, sender string) {
		defer wg.Done()
		val := 0
		for {
			select {
			case <-done:
				fmt.Println("Producer done")
				close(output)
				return
			case <-time.After(400 * time.Millisecond):
				output <- payload{sender: sender, value: val}
				val++
			}
		}
	}(wg, sender)
	return output
}

func consumer(producers []chan payload, fanIn chan<- payload, wg *sync.WaitGroup, counter *int64) {
	for _, producer := range producers {
		go func(fanIn chan<- payload, producer <-chan payload, wg *sync.WaitGroup, counter *int64) {
			defer wg.Done()
			for val := range producer {
				fmt.Println("Consume sender: ", val.sender, " value: ", val.value)
				fanIn <- val
			}
			atomic.AddInt64(counter, 1)
		}(fanIn, producer, wg, counter)
	}
}
