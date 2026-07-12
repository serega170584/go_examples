package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	res, err := adapter(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(*res)
}

func adapter(ctx context.Context) (*int, error) {
	ch := make(chan int, 1)
	go func() {
		ch <- slow()
		close(ch)
	}()

	select {
	case v := <-ch:
		return &v, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func slow() int {
	time.Sleep(3 * time.Second)
	return rand.Intn(10000)
}
