package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[int]int)
	wg := sync.WaitGroup{}
	wg.Add(1)
	mu := sync.Mutex{}
	go func(m map[int]int, wg *sync.WaitGroup, mu *sync.Mutex) {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			mu.Lock()
			fmt.Println(m[100])
			mu.Unlock()
		}
	}(m, &wg, &mu)

	for i := 0; i < 10000; i++ {
		mu.Lock()
		m[i] = i
		mu.Unlock()
	}

	wg.Wait()
}
