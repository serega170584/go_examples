package main

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	v, err := adapter(ctx)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	fmt.Println("Result " + strconv.Itoa(*v))
}

func getVal() *int {
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	v := rand.Intn(1000)
	return &v
}

func adapter(ctx context.Context) (*int, error) {
	in := make(chan *int)
	go func(in chan *int) {
		in <- getVal()
	}(in)

	select {
	case v := <-in:
		return v, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
