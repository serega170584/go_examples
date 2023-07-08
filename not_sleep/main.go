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
	notSleep(ctx)
}

func notSleep(ctx context.Context) {
	select {
	case <-time.After(time.Duration(rand.Intn(5)) * time.Second):
		fmt.Println(rand.Intn(1000))
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
