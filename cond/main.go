package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	data := "123"
	var wg sync.WaitGroup
	wg.Add(2)
	go listen(cond, &data, &wg)
	go listen(cond, &data, &wg)
	wg.Wait()
	wg.Add(3)
	go func(data *string) {
		defer wg.Done()
		cond.L.Lock()
		*data = "456"
		cond.Broadcast()
		cond.L.Unlock()
	}(&data)
	wg.Wait()
}

func listen(cond *sync.Cond, data *string, wg *sync.WaitGroup) {
	defer wg.Done()
	cond.L.Lock()
	wg.Done()
	cond.Wait()
	fmt.Println(*data)
	cond.L.Unlock()
}
