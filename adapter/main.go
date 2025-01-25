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

func adapter(ctx context.Context) (*int, error) {
	ch := make(chan int)
	go func() {
		v := something()
		ch <- v
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case v := <-ch:
		return &v, nil
	}
}

func something() int {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	return rand.Intn(10000)
}
