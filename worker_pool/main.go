package main

import (
	"fmt"
	"sync"
)

const wpSize = 6

func main() {
	a := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		a[i] = i
	}

	in := make(chan int)
	go func() {
		for _, v := range a {
			in <- v
		}
		close(in)
	}()

	var wg sync.WaitGroup
	wg.Add(wpSize)
	for i := 0; i < wpSize; i++ {
		go func() {
			defer wg.Done()
			for v := range in {
				fmt.Println(v)
			}
		}()
	}

	wg.Wait()
}
