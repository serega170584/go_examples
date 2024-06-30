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
	return rand.Intn(1000)
}

func adapter(ctx context.Context) (*int, error) {
	in := make(chan int)
	go func() {
		in <- something()
		close(in)
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case res := <-in:
		return &res, nil
	}
}
