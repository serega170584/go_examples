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
	ch := make(chan struct{})
	res, err := adapter(ctx, ch)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

func something() {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	fmt.Println("success")
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
