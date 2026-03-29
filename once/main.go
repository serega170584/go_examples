package main

import (
	"fmt"
	"sync"
)

func main() {
	done := make(chan struct{})
	c := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	wg.Add(2)
	data := 123
	go func() {
		c.L.Lock()
		runtime
		data = 456
		c.Broadcast()
		wg.Wait()
		c.L.Unlock()
		done <- struct{}{}
	}()
	go func() {
		defer wg.Done()
		c.L.Lock()
		c.Wait()
		fmt.Println(data)
		c.L.Unlock()
	}()
	go func() {
		defer wg.Done()
		c.L.Lock()
		c.Wait()
		fmt.Println(data)
		c.L.Unlock()
	}()
	<-done
}
