package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	cnt := 1000
	items := make([]int, cnt)
	existed := make(map[int]int, cnt)
	unique := make(chan int, cnt)

	for i := range items {
		items[i] = rand.Intn(20)
	}

	mu := sync.Mutex{}

	for _, double := range items {
		double := double
		go func() {
			mu.Lock()
			defer mu.Unlock()
			if _, ok := existed[double]; !ok {
				existed[double] = double
				unique <- double
			}

			cnt--
			if cnt == 0 {
				close(unique)
			}
		}()
	}

	for un := range unique {
		fmt.Println(un)
	}
}
