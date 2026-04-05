package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var mu sync.RWMutex
	go func() {
		time.Sleep(3 * time.Second)
		mu.RUnlock()
		fmt.Println("End 123")
	}()
	mu.RLock()
	fmt.Println("123")
	mu.RLock()
	fmt.Println("456")
	mu.RUnlock()

	t := time.Now()
	defer func() {
		fmt.Println(time.Since(t))
	}()

	var cnt uint64
	var wg sync.WaitGroup
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go func() {
			defer wg.Done()
			atomic.AddUint64(&cnt, 1)
		}()
	}

	wg.Wait()

	fmt.Println(cnt)

	var g errgroup.Group

	var urls = []string{
		"http://www.google1111.com",
		"http://www.google.com", // This will likely return an error
		"http://www.golang.org",
	}

	for _, url := range urls {
		g.Go(func() error {
			_, err := http.Get(url)
			return err
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Printf("An error occurred: %v\n", err)
		return
	}

	fmt.Println("Successfully fetched all URLs.")
}
