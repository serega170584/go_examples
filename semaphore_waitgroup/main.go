package main

import (
	"fmt"
	"math/rand"
	"time"
)

const N = 10

func main() {
	semaphore := make(chan int, N)
	results := make([]int, N)

	for i := 0; i < N; i++ {
		results[i] = i + 10
	}

	go func() {
		for i := 0; i < N; i++ {
			semaphore <- calculate(results[i])
		}
	}()

	for i := 0; i < N; i++ {
		fmt.Println("Got value ", <-semaphore)
	}
}

func calculate(val int) int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
	return val * 2
}
