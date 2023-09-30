package main

import (
	"fmt"
	"sync"
	"time"
)

type payload struct {
	name string
	num  int
}

func main() {
	wg := sync.WaitGroup{}
	cnt := 3
	producers := make([]chan payload, cnt)
	done := make(chan struct{})
	fanIn := make(chan payload)

	wg.Add(cnt)
	producers = append(producers, producer("Alice", &wg, done))
	producers = append(producers, producer("Mark", &wg, done))
	producers = append(producers, producer("Max", &wg, done))

	wg.Add(cnt)
	consumer(producers, &wg, done, fanIn)

	go func() {
		for val := range fanIn {
			fmt.Printf("fan in got %v\n", val)
		}
	}()

	time.Sleep(3 * time.Second)
	close(done)
	wg.Wait()
	close(fanIn)
}

func producer(name string, wg *sync.WaitGroup, done chan struct{}) chan payload {
	ch := make(chan payload)
	go func() {
		defer wg.Done()
		var i int
		for {
			select {
			case <-done:
				close(ch)
				return
			case <-time.After(500 * time.Millisecond):
				i++
				p := payload{name: name, num: i}
				ch <- p
			}
		}
	}()
	return ch
}

func consumer(producers []chan payload, wg *sync.WaitGroup, done chan struct{}, fanIn chan payload) {
	for i, producer := range producers {
		producer := producer
		i := i + 1
		go func() {
			defer wg.Done()
			for {
				select {
				case <-done:
					fmt.Println("Consume of producer", i, "completed")
					return
				case v := <-producer:
					fmt.Println("Consumer of producer", i, "got value", v.num, "from", v.name)
					fanIn <- v
				}
			}
		}()
	}
}
