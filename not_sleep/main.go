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
	Run(ctx)
}

func Run(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	case <-time.After(time.Duration(rand.Intn(5)) * time.Second):
		fmt.Println(rand.Intn(1000))
	}
}
