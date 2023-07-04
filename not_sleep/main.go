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
	res, err := Run(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

func Run(ctx context.Context) (int, error) {
	for {
		select {
		case <-ctx.Done():
			return 0, ctx.Err()
		case <-time.After(time.Duration(rand.Intn(5)) * time.Second):
			return rand.Intn(1000), nil
		}
	}
}
