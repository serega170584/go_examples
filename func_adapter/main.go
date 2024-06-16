package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	v, err := adapter(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*v)
	defer cancel()
}

func something() int {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	return rand.Intn(1000)
}

func adapter(ctx context.Context) (*int, error) {
	in := make(chan int)

	go func() {
		in <- something()
	}()

	select {
	case v := <-in:
		return &v, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
