package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	cnt := 1000
	doubles := make([]int, cnt)
	existed := make(map[int]int)
	unique := make(chan int, cnt)
	counter := make(chan struct{})

	for i := 0; i < cnt; i++ {
		doubles[i] = rand.Intn(10)
	}

	mu := sync.Mutex{}

	go func() {
	loop:
		for range counter {
			cnt--
			if cnt == 0 {
				break loop
			}
		}
		close(counter)
		close(unique)
	}()

	for _, val := range doubles {
		val := val
		go func() {
			mu.Lock()
			defer mu.Unlock()
			if _, ok := existed[val]; !ok {
				existed[val] = val
				unique <- val
			}
			counter <- struct{}{}
		}()
	}

	for val := range unique {
		fmt.Println(val)
	}
}
