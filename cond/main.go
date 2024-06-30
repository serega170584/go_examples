package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})
	wg := &sync.WaitGroup{}
	wg.Add(2)
	data := make(map[int]int, 0)
	go listen(c, data, wg)
	go listen(c, data, wg)
	go broadcast(c, data)
	wg.Wait()
}

func broadcast(c *sync.Cond, data map[int]int) {
	time.Sleep(time.Second)

	c.L.Lock()

	data[0] = 321

	c.Broadcast()

	fmt.Println("broadcast")

	c.L.Unlock()
}

func listen(c *sync.Cond, data map[int]int, wg *sync.WaitGroup) {
	defer wg.Done()
	c.L.Lock()

	c.Wait()

	fmt.Println("listen data ", data)

	c.L.Unlock()
}
