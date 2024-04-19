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

	v, err := adapter(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(*v)
}

func something() int {
	time.Sleep(time.Duration(rand.Intn(6)) * time.Second)
	return rand.Intn(100)
}

func adapter(ctx context.Context) (*int, error) {
	ch := make(chan int)

	go func() {
		ch <- something()
	}()

	select {
	case v := <-ch:
		return &v, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
