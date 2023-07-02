package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	cnt := 1000
	storage := make(map[int]int, cnt)
	unique := make(chan int, cnt)
	doubles := make([]int, cnt)

	for i := 0; i < cnt; i++ {
		doubles[i] = rand.Intn(20)
	}

	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(cnt)

	for _, val := range doubles {
		val := val
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			if _, ok := storage[val]; !ok {
				storage[val] = val
				unique <- val
			}
		}()
	}

	wg.Wait()

	close(unique)

	for val := range unique {
		fmt.Println(val)
	}
}
