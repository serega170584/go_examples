package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan struct{})
	mu := &sync.Mutex{}
	c := sync.NewCond(mu)
	go func(c *sync.Cond, ch chan struct{}) {
		c.L.Lock()
		c.Wait()
		fmt.Println("Ready")
		c.L.Unlock()
		ch <- struct{}{}
	}(c, ch)

	time.Sleep(time.Second)

	c.L.Lock()
	fmt.Println("Send")
	c.L.Unlock()
	c.Broadcast()
	fmt.Println("Send unlock")

	<-ch
}
