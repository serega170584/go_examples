package main

import (
	"fmt"
	"math/rand"
)

func main() {
	cnt := 1000
	doubles := make([]int, cnt)
	storage := make(map[int]struct{}, cnt)
	unique := make(chan int, cnt)
	for i := 0; i < cnt; i++ {
		doubles[i] = rand.Intn(20)
	}

	go func() {
		for _, double := range doubles {
			if _, ok := storage[double]; !ok {
				storage[double] = struct{}{}
				unique <- double
			}
		}
		close(unique)
	}()

	for val := range unique {
		fmt.Println(val)
	}
}
