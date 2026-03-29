package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(rand.Intn(5))*time.Second)
	defer cancel()

	v, err := adapter(ctx)

	if err != nil {
		fmt.Println("Error ", err)
		return
	}

	fmt.Printf("Success %d", *v)
}

func something() int {
	time.Sleep(3 * time.Second)
	res := rand.Intn(10000)
	return res
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
