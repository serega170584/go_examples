package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	done := make(chan struct{})

	wg := sync.WaitGroup{}
	wg.Add(2)

	res := generator(done, &wg)

	go func() {
		defer wg.Done()
		for val := range res {
			fmt.Println(val)
		}
	}()

	time.Sleep(3 * time.Second)
	close(done)
	wg.Wait()
}

func generator(done chan struct{}, wg *sync.WaitGroup) <-chan int {
	res := make(chan int)
	defer wg.Done()
	go func() {
		i := 0
		for {
			select {
			case <-done:
				close(res)
				return
			case <-time.After(500 * time.Millisecond):
				i++
				res <- i
			}
		}
	}()
	return res
}
