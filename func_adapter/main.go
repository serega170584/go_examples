package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	val, err := adapter(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
	cancel()
}

func something() {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
}

func adapter(ctx context.Context) (int, error) {
	ch := make(chan struct{})
	go func() {
		something()
		ch <- struct{}{}
	}()

	select {
	case <-ch:
		return rand.Intn(1000), nil
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}
