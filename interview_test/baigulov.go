package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	m, err := GetFiles(context.TODO(), "1", "2", "3", "4", "5")
	if err != nil {
		fmt.Println(time.Since(start))
		log.Fatalln(err)
	}

	fmt.Println(m)
	fmt.Println(time.Since(start))
}

type resultFromFile struct {
	name string
	data []byte
	err  error
}

// GetFiles пример функции, которую нужно оптимизировать.
func GetFiles(ctx context.Context, names ...string) (result map[string][]byte, err error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	if len(names) == 0 {
		return nil, nil
	}

	result = make(map[string][]byte, len(names))
	resultCh := make(chan resultFromFile)

	var wg sync.WaitGroup
	wg.Add(len(names))

	for _, name := range names {
		go func() {
			defer wg.Done()

			res, err := GetFile(ctx, name)
			resStruct := resultFromFile{name: name, data: res, err: err}
			select {
			case <-ctx.Done():
				fmt.Println("Context timed out")
			case resultCh <- resStruct:
				fmt.Println("Data Sent")
			}
		}()
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	for val := range resultCh {
		if val.err != nil {
			return nil, fmt.Errorf("Error %w", val.err)
		}
		result[val.name] = val.data
	}
	return result, nil
}

// GetFile является примером функции, которая относительно
// недолго выполняется при единичном вызове. Но достаточно
// долго если вызывать последовательно.
// Предположим, что оптимизировать в этой функции нечего.
func GetFile(ctx context.Context, name string) ([]byte, error) {
	if name == "" {
		return nil, fmt.Errorf("invalid name %q", name)
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-ticker.C:
	}

	if name == "invalid" {
		return nil, fmt.Errorf("invalid name %q", name)
	}

	b := make([]byte, 10)
	n, err := rand.Read(b)
	if err != nil {
		return nil, fmt.Errorf("getting file %q: %w", name, err)
	}

	return b[:n], nil
}
