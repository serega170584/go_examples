package main

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	v, err := adapter(ctx)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	fmt.Println("Value: " + strconv.Itoa(*v))
	defer cancel()
}

func adapter(ctx context.Context) (*int, error) {
	ch := make(chan int)
	go func() {
		res := something()
		ch <- res
	}()

	select {
	case v := <-ch:
		return &v, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func something() int {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Second)
	return rand.Intn(1000)
}
