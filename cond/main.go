package main

import (
	"fmt"
	"sync"
)

func main() {
	sync.Pool{}
	cond := sync.NewCond(&sync.Mutex{})
	data := make(map[string]string, 1)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go listen(data, cond, &wg)
	go listen(data, cond, &wg)
	wg.Wait()
	wg.Add(1)
	go func() {
		defer wg.Done()
		cond.L.Lock()
		data["123"] = "123"
		cond.Broadcast()
		wg.Add(2)
		cond.L.Unlock()
	}()
	wg.Wait()
}

func listen(data map[string]string, cond *sync.Cond, wg *sync.WaitGroup) {
	cond.L.Lock()
	wg.Done()
	cond.Wait()
	fmt.Println(data)
	cond.L.Unlock()
	wg.Done()
}
