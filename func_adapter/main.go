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
	resCh := make(chan struct{})
	res, err := adapter(ctx, resCh)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

func something() {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
}

func adapter(ctx context.Context, resCh chan struct{}) (int, error) {
	go func() {
		something()
		resCh <- struct{}{}
	}()

	select {
	case <-resCh:
		return rand.Intn(1000), nil
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}
