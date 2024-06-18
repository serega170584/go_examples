package main

import (
	"fmt"
	"sync"
	"time"
)

func listen(name string, data map[string]string, c *sync.Cond, wg *sync.WaitGroup) {
	defer wg.Done()
	c.L.Lock()
	c.Wait()

	fmt.Printf("[%s] %s\n", name, data["key"])

	c.L.Unlock()

	fmt.Printf("[%s] %s\n", name, data["key"])

}

func broadcast(name string, data map[string]string, c *sync.Cond) {
	time.Sleep(time.Second)

	c.L.Lock()

	data["key"] = "value"

	fmt.Printf("[%s] данные получены\n", name)

	c.Broadcast()
	c.L.Unlock()
}

func main() {
	data := map[string]string{}

	cond := sync.NewCond(&sync.Mutex{})

	wg := &sync.WaitGroup{}

	wg.Add(2)

	go listen("слушатель 1", data, cond, wg)
	go listen("слушатель 2", data, cond, wg)

	go broadcast("источник", data, cond)

	wg.Wait()

}
