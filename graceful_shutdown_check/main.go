package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	done := make(chan struct{})

	wg.Add(1)
	ch := createChan(&wg, done)

	wg.Add(1)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		for val := range ch {
			fmt.Println("Completed ", val)
		}
	}(ch, &wg)

	time.Sleep(time.Second)
	close(done)
	wg.Wait()
}

func createChan(wg *sync.WaitGroup, done <-chan struct{}) <-chan int {
	ch := make(chan int)

	go func() {
		defer wg.Done()
		i := 0
		for {
			select {
			case <-done:
				close(ch)
				fmt.Println("Chan done")
				return
			case <-time.After(499 * time.Millisecond):
				fmt.Println(i)
				ch <- i
				i++
			}
		}
	}()

	return ch
}
