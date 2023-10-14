package main

import (
	"fmt"
	"sync"
	"time"
)

type payload struct {
	name string
	val  int
}

// producer -> consumer -> fanIn
func main() {
	done := make(chan struct{})
	producers := make([]<-chan payload, 3)
	fanIn := make(chan payload)

	wg := sync.WaitGroup{}
	wg.Add(3) // для продюсеров
	producers[0] = producer(done, &wg, "Alice")
	producers[1] = producer(done, &wg, "Max")
	producers[2] = producer(done, &wg, "Nick")

	wg.Add(3) // для консумеров
	consumer(&wg, fanIn, producers)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range fanIn {
			fmt.Println("Out from: ", val.name, ", val:", val.val)
		}
	}()

	time.Sleep(3 * time.Second)
	close(done)
	close(fanIn)
	wg.Wait() // в самом конце чтобы отработали все горутины

}

func producer(done <-chan struct{}, wg *sync.WaitGroup, name string) <-chan payload {
	res := make(chan payload)

	go func() {
		defer wg.Done()
		var i int
		for {
			select {
			case <-done:
				close(res)
				return
			case <-time.After(500 * time.Millisecond):
				i++
				fmt.Println("Producing of: ", name, ", value:", i)
				res <- payload{name: name, val: i}
			}
		}
	}()

	return res
}

func consumer(wg *sync.WaitGroup, fanIn chan<- payload, producers []<-chan payload) {
	for _, producer := range producers {
		producer := producer
		go func() {
			defer wg.Done()
			for val := range producer {
				fmt.Println("Consume ", val.name, ", got value:", val.val)
				fanIn <- val
			}
		}()
	}
}
