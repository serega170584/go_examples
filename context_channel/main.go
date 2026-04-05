package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		select {
		case <-ctx.Done():
			fmt.Println("123")
		}
	}()

	go func() {
		defer wg.Done()
		select {
		case <-ctx.Done():
			fmt.Println("456")
		}
	}()

	wg.Wait()

}
