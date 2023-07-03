package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan struct{})
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	num, err := adapter(ctx, ch)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num)
}

func something() {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
}

func adapter(ctx context.Context, ch chan struct{}) (int, error) {
	go func() {
		something()
		ch <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case <-ch:
		return rand.Intn(1000), nil
	}
}
