package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})
	data := ""
	go func() {
		time.Sleep(time.Nanosecond)
		c.L.Lock()
		data = "123123123"
		c.Broadcast()
		c.L.Unlock()
	}()
	go func() {
		c.L.Lock()
		c.Wait()
		fmt.Println(data)
		c.L.Unlock()
	}()
	go func() {
		c.L.Lock()
		c.Wait()
		fmt.Println(data)
		c.L.Unlock()
	}()
	time.Sleep(3 * time.Second)

}
