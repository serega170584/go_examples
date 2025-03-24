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
		fmt.Println(err.Error())
		return
	}
	fmt.Println(*res)
}

func something() int {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	return rand.Intn(10000)
}

func adapter(ctx context.Context) (*int, error) {
	ch := make(chan int)

	go func() {
		ch <- something()
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case res := <-ch:
		return &res, nil
	}
}
